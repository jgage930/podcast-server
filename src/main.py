from fastapi import FastAPI


app = FastAPI()


@app.get('/')
def index():
    return {
        'msg': 'podcast-server v0.0.1'
    }

