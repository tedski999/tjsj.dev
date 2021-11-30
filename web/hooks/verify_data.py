#!/bin/python

import os
import json

print("=== Verifying data integrity ===")

with open("data.json") as f:
    data = json.load(f)
with open("tjsj.dev.json") as f:
    template = json.load(f)

# Check every post has a route
for post in data["posts"]:
    if post["url"] not in template["pages"]:
        raise Exception("Route to post '" + post["title"] + "' not found in template file!")
print("Posts OK")

# Check every project has a route
for project in data["projects"]:
    if project["url"] not in template["pages"]:
        raise Exception("Route to project '" + project["title"] + "' not found in template file!")
print("Projects OK")

# Check every page has a file
for route, file in template["pages"].items():
    if not os.path.isfile(file) and not os.path.isfile(os.path.join(os.environ["SITEGEN_DST"], file)):
        raise Exception("Expected file " + file + " for page " + route + " is missing!")
print("Routes OK")

print()
