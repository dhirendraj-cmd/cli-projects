import os
import yt_dlp
from pathlib import Path
from concurrent import futures
from itertools import repeat


def download_yt_video(video_url: str, file_path: str):

    if not os.path.exists(file_path):
        os.makedirs(file_path, exist_ok=True)

    ydl_opts = {
        "format": "bestvideo+bestaudio/best",
        "outtmpl": os.path.join(file_path, "%(title)s.%(ext)s"),
        "merge_output_format": "mp4",
        "quiet": True,
    }

    # for video_url in video_url_lst:
    with yt_dlp.YoutubeDL(ydl_opts) as out:
        out.download([video_url])
    print(f"Video downloaded successfully to {file_path}!")


video_url_lst: list[str] = []

while True:
    video_url = input("> ").strip()
    
    if video_url.lower() == "exit":
        break
    
    if video_url:
        video_url_lst.append(video_url)


print(video_url_lst)

file_path = os.path.join(Path.home(), "uploadtesting/src")
# file_path = "downloaded_videos"

with futures.ThreadPoolExecutor(max_workers=8) as out:
    out.map(download_yt_video, video_url_lst, repeat(file_path))

# download_yt_video(video_url_lst=video_url_lst, file_path=file_path)