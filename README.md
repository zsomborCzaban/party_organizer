# Party Organizer

<div id="header" align="center">
  <img src="https://user-images.githubusercontent.com/74038190/212750147-854a394f-fee9-4080-9770-78a4b7ece53f.gif" width="500">
<br><br><p>
  <strong>
    hey there
    <img src="https://media.giphy.com/media/hvRJCLFzcasrR4ia7z/giphy.gif" width="30px"/>
  </strong>
  <p/>
  <p>Welcome to the repository of my Party Organizer webiste<p/>
  <p>You can find a working demo at <a href="http://130.110.2.198/"> http://130.110.2.198/</a>. This was one of my earielst project and it has some flaws. I'm building other projects instead of perfecting this one, because theres more to learn I feel like. <p/>
</div>

---

## Demo instructions

I believe that the website is straight forward and doesnt need instructions, but there are some functions that only party owners can see that could be easily missed in a quick review.

My reccommendations:
 - Check out the sign up process.
 - Check out the "A regular party" with its **owners account**
 - 1. **username:** recruiter
 - 2. **password:** IWillHireZsombor0

---

## Technical Overview

Hey there!

Before checking out my code, please note that this was my first ever time working with go and react. Most of my code was written after 8 hours of work and under a time pressure (I used this for my thesis at my Bsc), so often quick and dirty solutions were choose for the sake of "feeling like making progress". This mostly reflect on my frontend code, and I hope I'll never do a project again with this mentality. While I think my backend code turned out mostly fine, in my frontend code i made lots of mistakes.

The application has a static frontend and an API for the backend with sqlite database (that could be swapped out). 

On the backend I followed clean architecture

On the frontend my first mistake was not following any architecture, which I deeply regret. Now I would follow MVVM.

---

## Running the app locally

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