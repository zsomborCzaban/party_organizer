# Project Name

## Preface

*(This section will contain the project preface. Content to be added later.)*

---

## Demo

*(This section will contain demo instructions and links. Content to be added later.)*

---

## Technical Overview

*(This section will provide a technical introduction and overview. Content to be added later.)*

---

## Installation Instructions

- Clone the repository

### Backend Setup

1. Download the latest version of **Go**: [https://go.dev/dl/](https://go.dev/dl/)
   * If the latest version doesn’t work, use **1.23.3**
2. Set the `$GOROOT` environment variable to the folder where Go is installed, and add the `/bin` directory of the Go installation to your `$PATH` environment variable if it hasn’t been added automatically.
3. Create an **AWS user** and an **AWS S3 bucket** if you don’t already have them. Documentation:

    * [Amazon S3 User Guide](https://docs.aws.amazon.com/AmazonS3)
4. Set up an **SMTP server** if you don’t have one. Documentation:

    * [Google Support](https://support.google.com)
5. Create a `.env` file in the `party_organizer/backend` directory, using `.env.example` as a template, and fill it with your credentials.

6. (optional): If you want a populated database to start with, rename `example.db` to `application.db` inside the `party_organizer/backend` directory.

7. Run the following command inside `party_organizer/backend`:

   ```bash
   go mod tidy
   ```
8. Start the backend with:

   ```bash
   go run main.go
   ```

---

### Frontend Setup

1. Download the latest version of **Node.js**: [https://nodejs.org/en/download](https://nodejs.org/en/download)

    * If the latest version doesn’t work, use:

        * **Node.js:** v20.11.0
        * **npm:** v10.2.4
2. Open a terminal in `party_organizer/frontend` and install dependencies:

   ```bash
   npm install
   ```
3. Create a `.env` file in the `party_organizer/frontend` directory, using `.env.example` as a template, and fill it with your credentials.

4. Start the frontend with:

   ```bash
   npm run dev
   ```

---