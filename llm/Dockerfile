FROM python:3.9-slim

# Install dependencies
RUN pip install flask transformers torch soundfile

# Copy the API server code into the container
WORKDIR /app
COPY llm_api.py .

# Expose the server port
EXPOSE 5000

# Start the API server
CMD ["python", "llm_api.py"]
