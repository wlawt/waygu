import discord
import requests
import os

from discord.ext import commands
from dotenv import load_dotenv

load_dotenv()
TOKEN = os.getenv("DISCORD_TOKEN")
bot = commands.Bot(command_prefix="$")

@bot.command(name="remind")
async def remind(ctx):
  r = requests.get("http://localhost:8000/remind")
  print(r.text)
  await ctx.channel.send("remind")

@bot.command(name="cpu")
async def cpu(ctx):
  r = requests.get("http://localhost:8000/cpu")
  print(r.text)
  await ctx.channel.send("task manager stats")
  
bot.run(TOKEN)