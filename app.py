from flask import Flask, render_template

app = Flask(__name__, template_folder="./template", static_folder="./template/static")


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/show")
def sow():
    return render_template("show.html")


if __name__ == "__main__":
    app.run(port=8099, debug=True)

