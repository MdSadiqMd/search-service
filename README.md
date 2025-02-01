# search-service

![image](https://github.com/user-attachments/assets/05c2498c-ed77-42d8-b958-0bf383ef8ca2)

Can scale upto 10 Million Records

Building Up:
Golang already have string.contains() method but that will take around ~2 sec for 600k docs

Thus we use Inverted Index Approach to search 
we have he text will process it and create Inverted Index from it 
we store how many palces do the word is present and keep them in an index as shown figure below
![image](https://github.com/user-attachments/assets/9cdbb523-934d-47d7-ad18-9de33f29a908)

Now for creating Index the process is as shown below
![image](https://github.com/user-attachments/assets/6805648c-c059-46d6-95ec-f258023586f8)
