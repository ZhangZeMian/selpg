# selpg

## 测试结果
---
### 第一类测试

1.$ ./selpg -s=1 -e=1 inputFile
    image 0
    
2.$ ./selpg -s=1 -e=1 inputFile
  image 1
  
3.$ env | ./selpg -s=1 -e=2 -l=2
  image2
  此处将env命令的输出重定向到selpg， 打印前两页，每页2行文字。

4.$ ./selpg -s=1 -e=2 -l=2 MainActivity.txt > receiverText.txt
  image3
  image4
  selpg将1-2页内容写到标准输出；shell将标准输出重定向到receiverText.txt中
 
5.$ ./selpg -s=5 -e=7 MainActivity.txt 2>errorFile.txt
  MainActivity.txt只有5页，所以会将第五页内容打印到屏幕；报错信息会被重定向给errorFile.txt
  imaga10
  image11
  
6.$ ./selpg -s=3 -e=1 MainActivity.txt > receiverText.txt 2>errorFile.txt 
  这个命令会报错，因为e>s，报错信息保存到errorFile.txt中
  image5

7.$ ./selpg -s=3 -e=1 MainActivity.txt > receiverText.txt 2>/dev/null
  错误信息被丢弃；标准输出重定向到receiverText.txt中

8.$ ./selpg -s=5 -e=7 MainActivity.txt >/dev/null
  文件只有5页，所以这里会打印第五页后报错。标准输出被舍弃，错误信息显示在屏幕。
  image6

9.$ ./selpg -s=1 -e=3 -l=2 MainActivity.txt | wc
  selpg的输出被重定向到wc中，wc命令显示接收的输入内容的行数、字数、字符数
  image7

10.$ ./selpg -s=5 -e=7 MainActivity.txt 2>errorFile.txt | wc
  MainActivity.txt只有5页，所以会将第五页内容重定向给wc；报错信息会被重定向给errorFile.txt
  image8
  image9

---
### 第二类测试
1.$ ./selpg -s=1 -e=3 -l=2 MainActivity.txt
  将MainActivity.txt的1-3页，页行数为2，打印到标准输出。
  image12

2.$ ./selpg -s=1 -e=2 -f MainActivity.txt
  以'\f'作为一页的结尾，在MainActivity.txt文件中，每2行行尾('\n'之后)都有一个'\f'，所以应打印出6行（由于'\f'在linux打印到屏幕时表现为换行，所以是4+2，如果保存到文件中则是4行）
  image13

3.$ ./selpg -s=1 -e=2 -f=3 -d=receiver MainActivity.txt
  receiver是我自己写的一个小程序，它会将接收到的标准输入打印到文件receiver.txt中。-d=receiver表示创建子进程receiver，通过管道将selpg的标准输出作为receiver的标准输入。
  image14
  
4.$ ./selpg -s=1 -e=3 MainActivity.txt > receiverText.txt 2>errorFile.txt &
  后台运行selpg程序。
  image15