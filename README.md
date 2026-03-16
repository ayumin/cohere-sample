# cohere-sample

Sample scripts demonstrating RAG (Retrieval-Augmented Generation) using the [Cohere API](https://cohere.com/) in Python, TypeScript, Ruby, and Go.

## Setup

Copy `.env.sample` to `.env` and set your Cohere API key:

```sh
cp .env.sample .env
```

Edit `.env`:

```
COHERE_API_KEY=your_api_key_here
```

---

## Python

**Requirements:** Python 3.8+

```sh
python -m venv .venv
source .venv/bin/activate
pip install cohere python-dotenv
python cohere-sample.py
```

---

## TypeScript

**Requirements:** Node.js 18+

```sh
npm install
npm start
```

---

## Ruby

**Requirements:** Ruby 3.0+ (rbenv recommended)

```sh
bundle install
ruby cohere-sample.rb
```

---

## Go

**Requirements:** Go 1.21+

```sh
go mod tidy
go run cohere-sample.go
```

---

## License

MIT
