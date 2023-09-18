# Create course

## Available Parameters

### Title

- It is the name of the course. This has not to be unique if you want to create more than one course with the same name.
- Maximum length is 60 characters
- It is required

### Language

- It is the language where the course is created
- It is a word of two characters representing the language. e.g norwegian => no
- It is required

### Description

- Optional description about the course.
- Maximum length is 200 characters

### SourceURL

Optional URL where the course comes from

### Tags

Optional array of tags to classify the course

### Image

Optional file image.

### Level

Mandatory paramter to classify the course as:

- Beginner 1 (A1)
- Beginner 2 (A2)
- Intermediate 1 (B1)
- Intermediate 2 (B2)
- Advanced 1 (C1)
- Advanced 2 (C2)

### Examples

`./go-lingq course create --title "er jeg liten?" --language no --level A1` -> this line will create a new course on the norwegian language, with the title "er jeg liten?" and a Beginner level (A1)

```shell
./go-lingq course create \
    --title "er jeg liten?" \
    --language no \
    --level A1 \
    --description "Det barnebok forklare forskjeller mellom stor og små" \
    --tags barnebok \
    --tags bok,barne \
    --image ./image.jpg
```
