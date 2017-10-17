# selpg
---
## 程序基本说明

本文演示如何用Go语言编写与 cat、ls、pr 和 mv 等标准命令类似的 Linux 命令行实用程序。我选择了一个名为 selpg 的实用程序，这个名称代表 SELect PaGes。selpg 允许用户指定从输入文本抽取的页的范围，这些输入文本可以来自文件/标准输入/另一个进程。

该程序是Go语言版本，具体程序说明可参考本程序的C语言版本,链接：[selpg程序说明][1]

---
## 作业亮点？

1.实现了-d选项，启用子进程并建立管道连接，数据能正确传递。

---
## 注意事项
1.为了测试-d选项，我创建了一个程序receiver.go，要运行selpg程序前需先编译receiver，并放在同个文件夹目录中。

---
### 
## 测试结果

### 第一类测试

1.$ ./selpg -s=1 -e=1 inputFile
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/0.png)
    
2.$ ./selpg -s=1 -e=1 inputFile
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/1.png)
  
3.$ env | ./selpg -s=1 -e=2 -l=2
  此处将env命令的输出重定向到selpg， 打印前两页，每页2行文字。
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/2.png)

4.$ ./selpg -s=1 -e=2 -l=2 MainActivity.txt > receiverText.txt
selpg将1-2页内容写到标准输出；shell将标准输出重定向到receiverText.txt中

  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/3.png)
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/4.png)
 
5.$ ./selpg -s=5 -e=7 MainActivity.txt 2>errorFile.txt
  MainActivity.txt只有5页，所以会将第五页内容打印到屏幕；报错信息会被重定向给errorFile.txt
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/10.png)
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/11.png)
  
6.$ ./selpg -s=3 -e=1 MainActivity.txt > receiverText.txt 2>errorFile.txt 
  这个命令会报错，因为e>s，报错信息保存到errorFile.txt中
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/5.png)

7.$ ./selpg -s=3 -e=1 MainActivity.txt > receiverText.txt 2>/dev/null
  错误信息被丢弃；标准输出重定向到receiverText.txt中

8.$ ./selpg -s=5 -e=7 MainActivity.txt >/dev/null
  文件只有5页，所以这里会打印第五页后报错。标准输出被舍弃，错误信息显示在屏幕。
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/6.png)

9.$ ./selpg -s=1 -e=3 -l=2 MainActivity.txt | wc
  selpg的输出被重定向到wc中，wc命令显示接收的输入内容的行数、字数、字符数
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/7.png)

10.$ ./selpg -s=5 -e=7 MainActivity.txt 2>errorFile.txt | wc
  MainActivity.txt只有5页，所以会将第五页内容重定向给wc；报错信息会被重定向给errorFile.txt
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/8.png)
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/9.png)

---
### 第二类测试
1.$ ./selpg -s=1 -e=3 -l=2 MainActivity.txt
  将MainActivity.txt的1-3页，页行数为2，打印到标准输出。
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/12.png)

2.$ ./selpg -s=1 -e=2 -f MainActivity.txt
  以'\f'作为一页的结尾，在MainActivity.txt文件中，每2行行尾('\n'之后)都有一个'\f'，所以应打印出6行（由于'\f'在linux打印到屏幕时表现为换行，所以是4+2，如果保存到文件中则是4行）
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/13.png)

3.$ ./selpg -s=1 -e=2 -f=3 -d=receiver MainActivity.txt
  receiver是我自己写的一个小程序，它会将接收到的标准输入打印到文件receiver.txt中。-d=receiver表示创建子进程receiver，通过管道将selpg的标准输出作为receiver的标准输入。
  
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/14.png)
  
4.$ ./selpg -s=1 -e=3 MainActivity.txt > receiverText.txt 2>errorFile.txt &
  后台运行selpg程序。
  **粗体文本**
  ![image](https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/15.png)


  [1]: https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html
  [2]: https://raw.githubusercontent.com/ZhangZekun/selpg/master/image/0.png