from flask import Flask, jsonify
def create_app():
    app = Flask(__name__)

    @app.route('/healthz')
    def health_check():
        health = {'status': 'ok'}
        return jsonify(health)

    @app.route('/version')
    def version():
        version = {'name': 'Kim JeongTae', 'version': '0.0.1'}
        return jsonify(version)
    
    return app

if __name__ == '__main__':
    app = create_app()
    app.run(host='0.0.0.0',port=8080)