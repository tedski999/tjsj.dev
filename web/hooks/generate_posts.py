#!/bin/python

import os
import json
import markdown

print("=== Generating posts ===")

with open("data.json") as f:
    data = json.load(f)
with open("tjsj.dev.json") as f:
    template = json.load(f)
with open("posts/template.html") as f:
    postTemplate = f.read()
postTemplateDate = os.path.getmtime("posts/template.html")

print("Scanning posts Markdown files for out-of-date HTML pages...")
md = markdown.Markdown()
for post in data["posts"]:

    # Don't regenerate post if template or markdown haven't changed
    if os.path.isfile(template["pages"][post["url"]]):
        srcDate = os.path.getmtime(post["markdown"])
        dstDate = os.path.getmtime(template["pages"][post["url"]])
        if dstDate >= postTemplateDate and dstDate >= srcDate:
            print("Skipping " + template["pages"][post["url"]])
            continue

    print("Generating " + template["pages"][post["url"]] + "...")
    with open(post["markdown"]) as f:
        text = f.read()
    html = postTemplate % (post["title"], post["description"], post["url"], post["title"], md.reset().convert(text))
    outfile = template["pages"][post["url"]]
    os.makedirs(os.path.dirname(outfile), exist_ok=True)
    with open(outfile, "w") as f:
        f.write(html)

print()
