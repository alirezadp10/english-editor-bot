# Save this as `whisper_api.py`
from transformers import pipeline
from flask import Flask, request, jsonify

app = Flask(__name__)

# Initialize the ASR pipeline
pipe = pipeline("automatic-speech-recognition", model="openai/whisper-large-v3")

@app.route("/transcribe", methods=["POST"])
def transcribe():
#     if "file" not in request.files:
#         return jsonify({"error": "No file uploaded"}), 400
#
#     audio_file = request.files["file"]
#     transcription = pipe(audio_file)["text"]

    return jsonify({"transcription": "hello world!"})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
