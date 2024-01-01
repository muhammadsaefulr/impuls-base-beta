import instaloader
from urllib.parse import unquote
import re
import sys
import json

if len(sys.argv) != 2:
    print("Usage: python script.py <instagram_post_url>")
    sys.exit(1)

url = sys.argv[1]
shortcode_pattern = r'/reel/([^/?]+)'
decodedurl = unquote(url)

match = re.search(shortcode_pattern, decodedurl)

if match:
    shortcodeurl = match.group(1)
    print("Shortcode:", shortcodeurl)
else:
    print("No shortcode found.")
    

loader = instaloader.Instaloader()

try:
    shortcode = shortcodeurl
    post = instaloader.Post.from_shortcode(loader.context, shortcode)
except Exception as e:
    print(f"Error: {e}")
    sys.exit(1)

username_url = post.owner_username
post_url = post.video_url if post.is_video else post.url
desc_url = post.caption
thumbnail_url = post.url

response = {
    'dataDetail': [
        {
            'username': username_url,
            'postUrl': post_url,
            'thumbnailUrl': thumbnail_url,
            'postDesc': desc_url
        }
    ]
}

output_file = 'result_instagram_download.json'
with open(output_file, 'w') as json_file:
    json.dump(response, json_file, indent=4)
# print(response['dataDetail'][0]['postUrl'])