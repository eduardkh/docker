from flask import Flask, request

# return remote request IP


def get_remote_ip():
    return request.remote_addr


# create a Flask app
app = Flask(__name__)

# define a route


@app.route('/')
def whatismyip():
    return get_remote_ip()


# run the app
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
