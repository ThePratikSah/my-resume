# My Resume

A forkable, version-controlled resume workflow powered by GitHub Actions, `md-to-pdf`, and AI assistants. Maintain your base resume in Markdown, and generate tailored PDFs per job application, all from your editor or terminal.

---

## How it works

The repo uses [md-to-pdf](https://www.npmjs.com/package/md-to-pdf) to convert `resume.md` into a styled `resume.pdf`. A GitHub Action runs on every push to rebuild the PDF automatically.

---

## Quick start

### 1. Fork this repo

Click the Fork button on GitHub, then clone your fork:

```
git clone https://github.com/<your-username>/my-resume.git
cd my-resume
npm install
```

### 2. Write your base resume

Edit `resume.md` with your details. The frontmatter controls PDF layout:

```yaml
---
stylesheet:
  - style.css
pdf_options:
  format: A4
  margin: 0mm
  printBackground: true
---
```

Customise `style.css` to match your personal branding.

### 3. Preview locally

```
npm run pdf
```

This generates `resume.pdf` in the project root.

### 4. Commit your base resume

```
git add resume.md resume.pdf style.css
git commit -m "feat: add base resume"
git push
```

The GitHub Action will rebuild `resume.pdf` on push. The commit-created PDF and the Action-built PDF should match.

---

## Tailoring for a job description

When applying to a specific role, create a branch for that company, ask an AI assistant to tailor the resume, and push.

### Step-by-step

1. **Create a branch**

   ```
   git checkout -b apply/<company-name>
   ```

2. **Ask an AI assistant to tailor the resume**

   Provide the assistant with:
   - Your current `resume.md`
   - The job description (paste it or give a link)

   Prompt example for Claude code, Codex or Opencode:

   ```
   Update resume.md to tailor it for this job description.
   Keep the same YAML frontmatter, Markdown structure, and
   general layout. Reword bullet points and skills to match
   the JD's keywords and requirements. Do not fabricate
   experience, only rephrase and reorder what is already there.
   ```

3. **Build and verify**

   ```
   npm run pdf
   ```

   Open `resume.pdf` to check the output.

4. **Commit and push**

   ```
   git add resume.md resume.pdf
   git commit -m "tailor resume for <company-name>"
   git push -u origin apply/<company-name>
   ```

   The GitHub Action will rebuild `resume.pdf` on push.

---

## Keeping your base resume in sync

When your career progresses, update `main` first, then rebase your application branches:

```
git checkout main
# edit resume.md with your new role/achievements
git add resume.md resume.pdf && git commit -m "update base resume"
git push

git checkout apply/<company-name>
git rebase main
# re-tailor if needed
git push --force-with-lease
```

---

## Resume

[View Resume PDF](resume.pdf)

The PDF is auto-generated from `resume.md` on every push via the [Build Resume PDF](.github/workflows/build.yml) workflow.
