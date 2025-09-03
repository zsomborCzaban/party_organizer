# Party Organizer

<div id="header" align="center">
  <img src="https://user-images.githubusercontent.com/74038190/212750147-854a394f-fee9-4080-9770-78a4b7ece53f.gif" width="500">
<br><br><p>
  <strong>
    hey there
    <img src="https://media.giphy.com/media/hvRJCLFzcasrR4ia7z/giphy.gif" width="30px"/>
  </strong>
  <p/>
  <p>Welcome to the repository for my Party Organizer website!</p>
  <p>You can try a live demo at <a href="http://130.110.2.198/">http://130.110.2.198/</a>. This was one of my earliest projects, therefore it has some flaws. Instead of perfecting it, I decided to move on and build other projects to learn more.</p>
</div>

---

## 📖 Demo Instructions

I believe the website is straightforward and doesn’t require much explanation. However, there are some features only **party owners** can access, which could be easy to miss during a quick review.

Here are my recommendations:
- Try out the sign-up process.
- Check out the **"A Regular Party"** page using the owner’s account:
    1. **Username:** recruiter
    2. **Password:** IWillHireZsombor0

---

## 🛠️ Technical Overview

Hey there!

Before diving into the code, keep in mind that this was my first time working with Go and React. Most of this project was built after long workdays and under time pressure (I used it for my BSc thesis), so I often chose quick and dirty solutions just to feel like I was making progress.

This is most obvious in the frontend code, which I now see has plenty of mistakes. I hope to never make the mistake of adding technical debt in the name of making progress.

The app consists of:
- A **static frontend**
- A **backend API** with a **SQLite** database (that can be swapped)

On the backend, I followed Clean Architecture. 
On the frontend, I made the mistake of not following any architecture, which I deeply regret (and would have regretted even more if I had to do unit test). If I did it again, I’d follow **MVVM** architecture.

---

## 🚀 Running the app locally

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