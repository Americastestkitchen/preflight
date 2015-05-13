# Preflight
Don't take off with only one engine

This will one day be a massive, concurrent, automated code review. For now, let's just exclude Ruby syntax errors from the human code review process.

What do do with this:

1. Download the executable from releases.
https://github.com/Americastestkitchen/preflight/releases
2. Copy the **preflight** executable to your project’s .git/hooks directory, renaming it "**pre-commit**"
ex. ```cp ~/Downloads/preflight your_project/.git/hooks/pre-commit```
3. Make it executable.  
```chmod a+x .git/hooks/pre-commit```
4. Try to commit some bogus Ruby. It'll yell at you.
