import os 
import requests
from flask import Flask, jsonify, request
import requests
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate

app = Flask(__name__)

basedir = os.path.abspath(os.path.dirname(__file__))
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///' + os.path.join(basedir, 'app.db')
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)
migrate = Migrate(app, db)

class Article(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(120), nullable=False)

@app.route('/api/articles', methods=['GET'])
def get_articles():
    articles = Article.query.all()
    articles_list = [{"id": article.id, "title": article.title} for article in articles]
    return jsonify(articles_list)

@app.route('/api/articles', methods=['POST'])
def add_articles():
    new_article_data = request.json
    new_article = Article(title=new_article_data['title'])
    db.session.add(new_article)
    db.session.commit()

    try:
        requests.post('http://localhost:8081/notify')
    except Exception as e:
        print(f"Erreur lors de l'envoi de la notification WebSocket: {e}")
    return jsonify({"id": new_article.id, "title": new_article.title}), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)