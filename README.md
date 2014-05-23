# Preflight
Don't take off with only one engine

This will one day be a massive, concurrent, automated code review. For now, let's just exclude Ruby syntax errors from the human code review process.

What do do with this:

1. Download the executable from releases.
2. Copy it to .git/hooks/pre-commit.
3. Make it executable (chmod a+x .git/hooks/pre-commit).
4. Try to commit some bogus Ruby. It'll yell at you.
