#!/bin/python

from feedgen.feed import FeedGenerator
import datetime
import json
import os

print("=== Generating feeds ===")

with open("data.json") as f:
    data = json.load(f)

# Create an Atom and RSS feed for posts
print("Generating posts feed...")
data["posts"].sort(key=lambda x: x["published"]["date"], reverse=True)
fg = FeedGenerator()
fg.title("tjsj.dev")
fg.description("Ted Johnson's computer science and game engine devblog")
fg.id("https://tjsj.dev/posts/")
fg.icon("https://tjsj.dev/icon/favicon-32x32.png")
fg.logo("https://tjsj.dev/img/logo.png")
fg.link(href="https://tjsj.dev/posts/", rel="alternate")
latestPostDateTime  = data["posts"][0]["published"]["date"] + " " + data["posts"][0]["published"]["time"] + " Z"
fg.lastBuildDate(datetime.datetime.strptime(latestPostDateTime, "%Y-%m-%d %H:%M:%S %z"))
fg.author(name="Ted Johnson", email="tedjohnsonjs@gmail.com")
fg.webMaster("tedjohnsonjs@gmail.com")
fg.language("en")
fg.generator("")

# Add 20 of the most recent posts
for post in data["posts"][:20]:
    fe = fg.add_entry()
    fe.title(post["title"])
    fe.summary(post["description"])
    postDateTime = post["published"]["date"] + " " + post["published"]["time"] + " Z"
    fe.published(datetime.datetime.strptime(postDateTime, "%Y-%m-%d %H:%M:%S %z"))
    postDateTime = post["updated"]["date"] + " " + post["updated"]["time"] + " Z"
    fe.updated(datetime.datetime.strptime(postDateTime, "%Y-%m-%d %H:%M:%S %z"))
    fe.guid("https://tjsj.dev" + post["url"], True)
    fe.link(href="https://tjsj.dev" + post["url"], rel="alternate")
    fe.link(href="https://tjsj.dev/posts.atom", rel="self")
    fe.author(name="Ted Johnson", email="tedjohnsonjs@gmail.com")

os.makedirs("public", exist_ok=True)
fg.atom_file("public/posts.atom")
fg.rss_file("public/posts.xml")

# Create an Atom and RSS feed for projects
print("Generating projects feed...")
data["projects"].sort(key=lambda x: x["published"]["date"], reverse=True)
fg = FeedGenerator()
fg.title("tjsj.dev")
fg.description("Ted Johnson's projects")
fg.id("https://tjsj.dev/projects/")
fg.icon("https://tjsj.dev/icon/favicon-32x32.png")
fg.logo("https://tjsj.dev/img/logo.png")
fg.link(href="https://tjsj.dev/projects/", rel="alternate")
latestProjectDateTime = data["projects"][0]["published"]["date"] + " " + data["projects"][0]["published"]["time"] + " Z"
fg.lastBuildDate(datetime.datetime.strptime(latestProjectDateTime, "%Y-%m-%d %H:%M:%S %z"))
fg.author(name="Ted Johnson", email="tedjohnsonjs@gmail.com")
fg.webMaster("tedjohnsonjs@gmail.com")
fg.language("en")
fg.generator("")

# Add 20 of the most recent projects
for project in data["projects"][:20]:
    fe = fg.add_entry()
    fe.title(project["title"])
    fe.summary(project["description"])
    projectDateTime = project["published"]["date"] + " " + project["published"]["time"] + " Z"
    fe.published(datetime.datetime.strptime(projectDateTime, "%Y-%m-%d %H:%M:%S %z"))
    projectDateTime = project["updated"]["date"] + " " + project["updated"]["time"] + " Z"
    fe.updated(datetime.datetime.strptime(projectDateTime, "%Y-%m-%d %H:%M:%S %z"))
    fe.guid("https://tjsj.dev" + project["url"], True)
    fe.link(href="https://tjsj.dev" + project["url"], rel="alternate")
    fe.link(href="https://tjsj.dev/projects.atom", rel="self")
    fe.author(name="Ted Johnson", email="tedjohnsonjs@gmail.com")

os.makedirs("public", exist_ok=True)
fg.atom_file("public/projects.atom")
fg.rss_file("public/projects.xml")

print()
