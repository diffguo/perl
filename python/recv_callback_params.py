# -*- coding: UTF-8 -*-

import re, sys, urllib, urllib2, urlparse
from BaseHTTPServer import HTTPServer, BaseHTTPRequestHandler

callback=""

class HTTPHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        mpath,margs=urllib.splitquery(self.path)
        self.do_action(mpath, margs)

        buf = 'It Works'
        
        self.protocal_version = "HTTP/1.1"

        self.send_response(200)

        self.send_header("Welcome", "Contect")       

        self.end_headers()

        self.wfile.write(buf)

    def do_action(self, path, args):
        import os
        if args:
            if path == callback:
                sys.stderr.write(args)
                os._exit(0)

def start_server(ip, port):
    http_server = HTTPServer((ip, int(port)), HTTPHandler)
    http_server.serve_forever()

def stop_server(server):
    server.sorket.close()

def get_public_ip():
    def visit(url):
        opener = urllib2.urlopen(url)
        if url == opener.geturl():
            str = opener.read()
            print str
        return re.search('\d+\.\d+\.\d+\.\d+',str).group(0)

    try:
        myip = visit("http://www.whereismyip.com/")
    except:
        import socket
        try:
            myip = socket.gethostbyname(socket.gethostname())
        except:
            myip = "127.0.0.1"

    return myip

def get_local_ip():
    import socket
    try:
        myip = socket.gethostbyname(socket.gethostname())
    except:
        myip = "127.0.0.1"

    return myip

def usage():
    print sys.argv[0],"\n-h help\n--ip listen_ip\n--port/-p listen port\n--callback callback url"

if __name__ == "__main__":
    import sys, getopt

    ip=""
    port=""

    opts, args = getopt.getopt(sys.argv[1:], "hp:", ["ip=", "port=", "callback="])
    for op, value in opts:
        if op in ("-p", "--port"):
            port = value
        elif op == "--ip":
            ip = value
        elif op == "--callback":
            callback = "/" + value
        elif op == "-h":
            usage()
            sys.exit()
        else:
            usage()
            sys.exit()

    if not ip:
        ip = get_local_ip()

    if not port:
        port = "8000"

    print "start http server ip: ", ip, " port: ", port

    start_server(ip, port)
