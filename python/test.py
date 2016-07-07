import requests, base64, json

def get_access_token():
    client_key = "3d1602888d5111e29f42782bcb058632"
    client_secret = "3d1606d48d5111e29f42782bcb058632"
    params = dict(

        client_id="3d1602888d5111e29f42782bcb058632",
        grant_type="authorization_code",
        code="902245924360c3eccea944ca7e8cd716",
        redirect_uri="http://localhost:9527/authenticate",
        scope="user,sports"
    )

    headers = {'Authorization': 'Basic %s' % base64.b64encode(client_key + ":" + client_secret)}
    response = requests.post("https://openapi.codoon.com/token",data=params, headers=headers)
    print response.text

    '''
    {"user_id":"618478a3-039a-483d-9863-796b1bac65b6","access_token":"782ddb08f3732eab1c0dabf1dbb674c0","token_type":"bearer","scope":"user sports","expire_in":93312000,"refresh_token":"e0b1220982d1f43cd89b86e9027b6cfc"}
    '''

def refresh_token():

    client_key = "3d1602888d5111e29f42782bcb058632"
    client_secret = "3d1606d48d5111e29f42782bcb058632"

    params = dict(
        client_id=client_key,
        grant_type="refresh_token",
        refresh_token="6aa25cb6e590e47730cb0eb3c303b8b8",
        scope="user,sports"
    )

    headers = {'Authorization': 'Basic %s' % base64.b64encode(client_key + ":" + client_secret)}
    response = requests.post("https://openapi.codoon.com/token", data=params, headers=headers)
    print response.text

def call_api_demo():

    headers = {"Authorization": "Bearer %s" % "782ddb08f3732eab1c0dabf1dbb674c0"}
    post_data = {}
    response = requests.request('POST', "https://openapi.codoon.com/api/verify_credentials", data=json.dumps(post_data), headers=headers)
    print response.content

if __name__ == "__main__":
    refresh_token()