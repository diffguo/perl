package data_box

import (
	"fmt"
	"testing"
)

type Obj struct {
	Id int
	Age int
	Name string
}

func TestStep1_Fill_DataSlice_WITH_OBJ_By_OBJ(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member Obj // With OBJ
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}
	err := NewDataBox(data).KeyField("Id").JoinByMap(func(keywords interface{}) interface{} {
		members := make(map[int]Obj) // By OBJ
		aids := keywords.([]int)
		for _, aid := range aids {
			members[aid] = Obj{Id: aid * 100, Name: "b100", Age:10}
		}

		return members
	}).SaveToField("Member")

	if err != nil {
		t.Fatal(err.Error())
	}

	// 打印
	for _, elem := range data {
		fmt.Printf("%+v\n", elem)
	}
}

func TestStep2_Fill_DataSlice_WITH_OBJ_By_OBJPTR(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member Obj // With OBJ
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}
	err := NewDataBox(data).KeyField("Id").JoinByMap(func(keywords interface{}) interface{} {
		members := make(map[int]*Obj) // By OBJ PTR
		aids := keywords.([]int)
		for _, aid := range aids {
			members[aid] = &Obj{Id: aid * 100, Name: "b100", Age:10}
		}

		return members
	}).SaveToField("Member")

	if err != nil {
		t.Fatal(err.Error())
	}

	//打印
	for _, elem := range data {
		fmt.Printf("%+v\n", elem)
	}
}

func TestStep3_Fill_DataSlice_WITH_OBJPTR_By_OBJPTR(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member *Obj // With OBJ PTR
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}
	err := NewDataBox(data).KeyField("Id").JoinByMap(func(keywords interface{}) interface{} {
		members := make(map[int]*Obj) // By OBJ PTR
		aids := keywords.([]int)
		for _, aid := range aids {
			members[aid] = &Obj{Id: aid * 100, Name: "b100", Age:10}
		}

		return members
	}).SaveToField("Member")

	if err != nil {
		t.Fatal(err.Error())
	}

	//打印
	for _, elem := range data {
		fmt.Printf("%+v %+v\n", elem, elem.Member)
	}
}

func TestStep4_Fill_DataSlice_WITH_OBJPTR_By_OBJ(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member *Obj // With OBJ PTR
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}

	//打印
	for _, elem := range data {
		fmt.Printf("%+v %+v\n", elem, elem.Member)
	}

	// No!!! 你不能这么用！！！
}

func TestStep5_Fill_DataSlice_WITH_OBJSlice_By_OBJSlice(t *testing.T) { // 与上面的例子同样，可以兼容Slice指针转换
	type T struct {
		Id int
		Name string
		Member []Obj // With []OBJ
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}
	err := NewDataBox(data).KeyField("Id").JoinByMap(func(keywords interface{}) interface{} {
		members := make(map[int][]Obj) // By []OBJ
		aids := keywords.([]int)
		for _, aid := range aids {
			members[aid] = []Obj{Obj{Id: aid * 100, Name: "b100", Age:10},Obj{Id: aid * 100, Name: "b101", Age:11}}
		}

		return members
	}).SaveToField("Member")

	if err != nil {
		t.Fatal(err.Error())
	}

	// 打印
	for _, elem := range data {
		fmt.Printf("%+v\n", elem)
	}
}

type Group struct {
	GroupId   int
	GroupName string
	Members   []Member
}

// 测试前向关联和后向关联
type Member struct {
	MemberId   int
	GroupId    int //前向关联
	UserId     string //后向关联
	MemberName string
	MemberInfo UserInfo
}

func FetchMembersByGroupIds(groupIds []int) []Member {
	var m []Member
	for _, gid := range groupIds {
		mid := gid * 100
		m = append(m, Member{MemberId: mid, GroupId: gid, MemberName: fmt.Sprintf("mname%d", mid), MemberInfo: UserInfo{}, UserId: "111-111"})
		m = append(m, Member{MemberId: mid + 1, GroupId: gid, MemberName: fmt.Sprintf("mname%d", mid+1), MemberInfo: UserInfo{}, UserId: "111-222"})
	}

	return m
}

type UserInfo struct {
	UserId   string
	UserName string
}

func FetchUserInfosByUserIds(userIds []string) []UserInfo {
	var u []UserInfo
	for _, uid := range userIds {
		u = append(u, UserInfo{UserId: uid, UserName: fmt.Sprintf("uname_for_%s", uid)})
	}

	return u
}

// 测试嵌套填充
func TestNestingFill(t *testing.T) {
	group := []Group{Group{GroupId: 1, GroupName: "gname1"},Group{GroupId: 2, GroupName: "gname2"}}
	err := NewDataBox(group).KeyField("GroupId").JoinByMap(func(keywords interface{}) interface{} {
		groupids := keywords.([]int)
		members := FetchMembersByGroupIds(groupids)

		err := NewDataBox(members).KeyField("UserId").JoinByMap(func(keywords interface{}) interface{} {
			userIds := keywords.([]string)
			userInfos := FetchUserInfosByUserIds(userIds)

			mu := make(map[string]*UserInfo)
			for i := 0; i < len(userInfos); i++ {
				mu[userInfos[i].UserId] = &userInfos[i]
			}

			return mu
		}).SaveToField("MemberInfo")

		if err != nil {
			t.Fatal(err.Error())
		}

		mm := make(map[int][]Member)
		for j := 0; j < len(members); j++ {
			if v, ok := mm[members[j].GroupId]; ok {
				v = append(v, members[j])
				mm[members[j].GroupId] = v
			} else {
				mm[members[j].GroupId] = []Member{members[j]}
			}
		}

		return mm
	}).SaveToField("Members")

	if err != nil {
		t.Fatal(err.Error())
	}

	//打印
	for _,g := range group {
		fmt.Printf("%+v\n", g)
	}
}

type ObjWithFather struct {
	Id int
	FatherId int
	Name string
}

func TestStep10_Fill_DataSlice_WITH_ObjSlice(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member *ObjWithFather // With OBJ (测试的时候可以换成Struct或者指针)
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}
	err := NewDataBox(data).KeyField("Id").JoinByObjs(func(keywords interface{}) interface{} {
		members := make([]*ObjWithFather, 0) // By OBJ (测试的时候可以换成Struct或者指针)
		aids := keywords.([]int)
		for _, aid := range aids {
			members = append(members, &ObjWithFather{Id: aid * 100, FatherId:aid, Name: "b100"})
		}

		return members
	}, "FatherId").SaveToField("Member")

	if err != nil {
		t.Fatal(err.Error())
	}

	// 打印
	for _, elem := range data {
		fmt.Printf("%+v  %+v\n", elem, elem.Member)
	}
}

func TestStep11_Fill_DataSlice_WITH_ObjSlice(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member []ObjWithFather // With OBJ. Not Support Member *[]ObjWithFather. Maybe: []ObjWithFather or []*ObjWithFather
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}
	err := NewDataBox(data).KeyField("Id").JoinByObjs(func(keywords interface{}) interface{} {
		members := make([]ObjWithFather, 0) // By OBJ Or By OBJPtr
		aids := keywords.([]int)
		for _, aid := range aids {
			members = append(members, ObjWithFather{Id: aid * 100, FatherId:aid, Name: "b100"})
		}

		return members
	}, "FatherId").SaveToField("Member")

	if err != nil {
		t.Fatal(err.Error())
	}

	// 打印
	for _, elem := range data {
		fmt.Printf("%+v\n", elem)
	}
}

func TestStep12_Fill_DataSlice_Continuely(t *testing.T) {
	type T struct {
		Id int
		Name string
		Member1 Obj // With OBJ
		Member2 Obj // With OBJ
	}

	data := []T{T{Id: 1, Name: "name1"}, T{Id: 2, Name: "name2"}}

	db := NewDataBox(data)

	err := db.KeyField("Id").JoinByMap(func(keywords interface{}) interface{} {
		members := make(map[int]Obj) // By OBJ
		aids := keywords.([]int)
		for _, aid := range aids {
			members[aid] = Obj{Id: aid * 100, Name: "b100", Age:10}
		}

		return members
	}).SaveToField("Member1")

	if err != nil {
		t.Fatal(err.Error())
	}

	err = db.JoinByMap(func(keywords interface{}) interface{} {
		members := make(map[int]Obj) // By OBJ
		aids := keywords.([]int)
		for _, aid := range aids {
			members[aid] = Obj{Id: aid * 1000, Name: "b100", Age:100}
		}

		return members
	}).SaveToField("Member2")

	if err != nil {
		t.Fatal(err.Error())
	}

	// 打印
	for _, elem := range data {
		fmt.Printf("%+v\n", elem)
	}
}