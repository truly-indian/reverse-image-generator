# Reverse Image Generator
- It is a backend service that powers a website which requires a reverse image search functionality.

# Tech Stack Used
    Golang
    Gin
    Google Wire

# API Details

- Reverse Image Genrator API:

      end_point: /api/v1/reverse_image_generator
      request_body: {
        "imageUrl": "some image url pointing to image",
        "pageToken": 1
      }
      response: {
        "Products": [
          {
            "productName": "name of the product",
            "price": "price of the product",
            "rating": "1"
          }
        ]
      }

# Steps To Run the Project
- Build from source

      Clone the repository.
         - Make sure that you have go, wire installed on your system if not please install them first.
         - Run this command to build
          make compile
         - Run this command to run the project
          ./out/reverse-image-generator start

- Running Dockerized Application

      - Pull the image using this command
          docker pull deepakmalik1999/reverse-image-generator:1.1.0
      - Run this command
          docker run -d \
              -e SERP_API_KEY=your_serp_api_key \
              -e OPEN_AI_API_KEY=your_apoen_ai_key \
              -e GROQ_AI_KEY=your groq_api_key \
              -p 8080:8080 \
              --name reverse-image-generator \
              deepakmalik1999/reverse-image-generator:1.1.0
  
# License
- MIT LICENSED

       MIT License

        Copyright (c) 2024 Deepak Kumar Malik
    
        Permission is hereby granted, free of charge, to any person obtaining a copy
        of this software and associated documentation files (the "Software"), to deal
        in the Software without restriction, including without limitation the rights
        to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
        copies of the Software, and to permit persons to whom the Software is
        furnished to do so, subject to the following conditions:
    
        The above copyright notice and this permission notice shall be included in all
        copies or substantial portions of the Software.
    
        THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
        IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
        FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
        AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
        LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
        OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
        SOFTWARE.


# Screen Shots
<img width="1415" alt="Screenshot 2024-08-27 at 6 54 18 PM" src="https://github.com/user-attachments/assets/3f8eca13-ce41-4013-a17b-9744ad0bf6fd">
