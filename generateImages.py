
from PIL import Image, ImageDraw, ImageFont

width, height = 800, 480  # Set the dimensions of your image
background_color = (255, 255, 255)  # RGB color for the background

blackImage = Image.new("RGB", (width, height), background_color)
blackDraw = ImageDraw.Draw(blackImage)

text = "Hello, World!"
text_color = (0, 0, 0)  # RGB color for the text
font_size = 36
font = ImageFont.truetype("arial.ttf", font_size)  # You need to specify a valid font file

text_position = (100, 100)  # X and Y coordinates for the text
blackDraw.text(text_position, text, fill=text_color, font=font)

redImage = Image.new("RGB", (width, height), background_color)
line_color = (0, 0, 0)  # RGB color for the lines
line_width = 20

redDraw = ImageDraw.Draw(redImage)
# Draw a line from (x1, y1) to (x2, y2)
redDraw.line([(0, 0), (800, 0)], fill=line_color, width=line_width)
redDraw.line([(0, 480), (800, 480)], fill=line_color, width=line_width)

redDraw.line([(0, 0), (0, 480)], fill=line_color, width=line_width)
redDraw.line([(800, 0), (800, 480)], fill=line_color, width=line_width)
blackImage2 = blackImage.convert('P', palette=Image.Palette.ADAPTIVE, colors=1)
redImage2 = redImage.convert('P', palette=Image.Palette.ADAPTIVE, colors=1)

blackImage2.save("output/black.bmp")
redImage2.save("output/red.bmp")
