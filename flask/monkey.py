from flask import Flask

app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Five Little Monkeys'

@app.route('/status/')
def get_status():
    return 'Five Little Monkeys are jumping on the bed'

@app.route('/jump/')
def jump_on_bed():
    return 'Monkey Bounced!'

@app.route('/fall/')
def fall_off_bed():
    return 'Monkey fell off the bed! Oh No!'

if __name__ == '__main__':
    app.run()
