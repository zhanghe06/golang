## Golang 下载

国内地址:
<http://golangtc.com/download>

官网英文地址(自备梯子):
<https://golang.org/dl/>

linux：
```
$ wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz
```

## Golang 官网

<http://golang.org/>


## Golang 文档

中文翻译:
<http://zh-golang.appsp0t.com/>

启动本机文档:
```
$ godoc -http=:6060
```
浏览器访问:
<http://localhost:6060>


## 设置系统环境变量

```
$ sudo vim /etc/profile
    # Golang environment variable
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/work
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
$ source /etc/profile
```

或者仅为当前用户设置环境变量
```
$ vim ~/.bashrc
$ source ~/.bashrc
```

注意：使用 zsh 的用户, 需要为 zsh 设置环境变量
```
$ vim ~/.zshrc
$ source ~/.zshrc
```


## 编译go源码
```
$ cd $GOROOT/src
$ ./all.bash
```
编译过程大概10分钟


## sublime 安装 Package Control

按下快捷键 Ctrl + `
输入以下内容,回车:
```
import urllib2,os; pf='Package Control.sublime-package'; ipp=sublime.installed_packages_path(); os.makedirs(ipp) if not os.path.exists(ipp) else None; urllib2.install_opener(urllib2.build_opener(urllib2.ProxyHandler())); open(os.path.join(ipp,pf),'wb').write(urllib2.urlopen('http://sublime.wbond.net/'+pf.replace(' ','%20')).read()); print 'Please restart Sublime Text to finish installation'
```
Preferences菜单下，多出一个菜单项 Package Control


## 安装GoSublime插件

Preferences >> Package Control

输入并选择 Install Package

再新输入框输入并选择 GoSublime


## GoSublime环境变量设置

Preferances >> Package Settings >> GoSublime >> Settings-Default
```
"env": {
    "GOPATH": "$HOME/code/golang",
    "GOROOT": "$HOME/go"
}
```
设置GoSublime缩进
```
// whether or not to indent with tabs (alignment is always done using spaces)
"fmt_tab_indent": true,

// the assumed width of the tab character (or number of spaces to indent with)
"fmt_tab_width": 4,
```

## 设置Sublime缩进

Preferances >> Settings-Default
```
// The number of spaces a tab is considered equal to
"tab_size": 4,

// Set to true to insert spaces when tab is pressed
"translate_tabs_to_spaces": true,
```

## 设置等宽字体

不设置的话,缩进看起来会很奇怪

Preferances >> Settings-User
```
"font_face": "Courier New",
"font_size": 11.0,
```


## Sublime 安装
下载地址：
<http://www.sublimetext.com/2> 官网下载 Sublime Text 2.0.2.tar.bz2

进入下载目录
```
$ cd 下载
```

解压
```
$ sudo tar -xvf Sublime\ Text\ 2.0.2.tar.bz2
$ ls
```

将解压后的sublime文件夹移动至系统文件夹
```
$ sudo mv Sublime\ Text\ 2 /usr/local/lib/
```

建立快捷方式
```
$ sudo ln -s /usr/local/lib/Sublime\ Text\ 2/sublime_text /usr/bin/subl
```


## Sublime 破解

首先查看一下sublime-text 2安装路径
```
$ whereis Sublime\ Text\ 2
Sublime Text 2: /usr/local/lib/Sublime Text 2
$ cd /usr/local/lib/Sublime\ Text\ 2
$ ls
Icon  lib  PackageSetup.py  PackageSetup.pyc  Pristine Packages  sublime_plugin.py  sublime_plugin.pyc  sublime_text
```
然后用vim打开sublime_text文件
```
$ vim sublime_text
```
然后用xxd把文件转换成十六进制格式:
```
:%!xxd
```
在vim中定位至“Thanks”文字附近:
```
/Thanks
```
接着查找数字串“3342”:
```
/3342
```
找到一处3342的地方大致是这个样子 ……4333 3342 3032…….

将这里的3342 改为3242,在vim中输入:
```
:s/3342/3242
```
将文件转换回去。
```
:%!xxd -r
```
保存文件、退出:
```
:wq
```
打开Sublime

help >> enter licence

将以下Licence贴进去
```
--BEGIN LICENSE--
China
Unlimited User License
EA7E-2861
BE67D2175D3569FDAB9EB5340FAD2822
E7B56B3397A76AA9FBE8AC3D3C65918B
DFC28F2EA158140D9E07853D594818EB
3A237B2E8E98ED257C269548F50EDA34
EF0C7F72D8917DB538A0245E46BFD6B1
85F4EDE331F253530ED67A5C19E92399
04C5F4A1AF4AF3DB5EC49C1FEE17CA76
7E369F8AAE4AC6C6E756B5882E1608B9
--END LICENSE--
```

提示：Thanks for purchasing!
成功！


## Sublime 桌面快捷方式
切换到root账户，新建sublime的快捷方式文件
```
$ sudo -i
# cd /usr/share/applications
# touch sublime.desktop
# gedit sublime.desktop
```
输入以下内容，保存
```
[Desktop Entry]
Name=Sublime Text 2
Comment=Sublime Text 2
Exec=/usr/bin/subl %F
Icon=/usr/local/lib/Sublime Text 2/Icon/256x256/sublime_text.png
Terminal=false
Type=Application
Categories=Application;Development;
StartupNotify=true
```


## Sublime 在ubuntu下的中文输入

以下步骤的前提是你已经设置了sublime的快捷方式

### 第一步：找到合适的输入法（fcitx框架）
```
$ apt-cache search fcitx
$ sudo apt-get install fcitx-pinyin
```
搜狗输入法就是fcitx框架，也可以直接下载linux版本的输入法

系统设置 >> 语言支持 >> 键盘输入方式系统
IBus 换成 fcitx
重启系统

### 第二步：编译动态链接库
保存下面的代码为sublime_imfix.c
```
/*
sublime-imfix.c
Use LD_PRELOAD to interpose some function to fix sublime input method support for linux.
By Cjacker Huang <jianzhong.huang at i-soft.com.cn>

gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPIC
LD_PRELOAD=./libsublime-imfix.so sublime_text
*/
#include <gtk/gtk.h>
#include <gdk/gdkx.h>
typedef GdkSegment GdkRegionBox;

struct _GdkRegion
{
    long size;
    long numRects;
    GdkRegionBox *rects;
    GdkRegionBox extents;
};

GtkIMContext *local_context;

void
gdk_region_get_clipbox (const GdkRegion *region,
            GdkRectangle    *rectangle)
{
    g_return_if_fail (region != NULL);
    g_return_if_fail (rectangle != NULL);

    rectangle->x = region->extents.x1;
    rectangle->y = region->extents.y1;
    rectangle->width = region->extents.x2 - region->extents.x1;
    rectangle->height = region->extents.y2 - region->extents.y1;
    GdkRectangle rect;
    rect.x = rectangle->x;
    rect.y = rectangle->y;
    rect.width = 0;
    rect.height = rectangle->height; 
    //The caret width is 2; 
    //Maybe sometimes we will make a mistake, but for most of the time, it should be the caret.
    if(rectangle->width == 2 && GTK_IS_IM_CONTEXT(local_context)) {
        gtk_im_context_set_cursor_location(local_context, rectangle);
  }
}

//this is needed, for example, if you input something in file dialog and return back the edit area
//context will lost, so here we set it again.

static GdkFilterReturn event_filter (GdkXEvent *xevent, GdkEvent *event, gpointer im_context)
{
    XEvent *xev = (XEvent *)xevent;
    if(xev->type == KeyRelease && GTK_IS_IM_CONTEXT(im_context)) {
        GdkWindow * win = g_object_get_data(G_OBJECT(im_context),"window");
        if(GDK_IS_WINDOW(win))
            gtk_im_context_set_client_window(im_context, win);
    }
    return GDK_FILTER_CONTINUE;
}

void gtk_im_context_set_client_window (GtkIMContext *context,
          GdkWindow    *window)
{
    GtkIMContextClass *klass;
    g_return_if_fail (GTK_IS_IM_CONTEXT (context));
    klass = GTK_IM_CONTEXT_GET_CLASS (context);
    if (klass->set_client_window)
        klass->set_client_window (context, window);

    if(!GDK_IS_WINDOW (window))
        return;
    g_object_set_data(G_OBJECT(context),"window",window);
    int width = gdk_window_get_width(window);
    int height = gdk_window_get_height(window);
    if(width != 0 && height !=0) {
        gtk_im_context_focus_in(context);
        local_context = context;
    }
    gdk_window_add_filter (window, event_filter, context); 
}
```

编译所需依赖：
```
$ sudo apt-get install build-essential
$ sudo apt-get install libgtk2.0-dev
```
编译动态库：
```
$ gcc -shared -o libsublime-imfix.so sublime_imfix.c `pkg-config --libs --cflags gtk+-2.0` -fPIC
```
编译完成后会得到这样的一个文件
libsublime-imfix.so

将编译好的文件放到/usr/lib目录下
```
$ sudo cp libsublime-imfix.so /usr/lib
```

### 第三步：为sublime的启动添加LD_PRELOAD环境变量
```
$ sudo -i
# cd /usr/share/applications
# gedit sublime.desktop
```
```
Exec=/usr/bin/subl %F
```
修改为：
```
Exec= bash -c 'LD_PRELOAD=/usr/lib/libsublime-imfix.so /usr/bin/subl' %F
```
保存，重新用快捷方式打开sublime

### 第四步：解决终端打开subl不能输入中文法的问题

以上步骤完成，通过快捷方式打开可以输入中文了，但是终端打开则还是不能输入

设置软链，保证终端启动前导入环境变量
```
$ alias subl='export LD_PRELOAD=/usr/lib/libsublime-imfix.so; /usr/bin/subl'
```

让　alias 永久生效
```
$ vim .bashrc
```
添加
```
alias subl='export LD_PRELOAD=/usr/lib/libsublime-imfix.so; /usr/bin/subl &>/dev/null'
```
```
$ source .bashrc
```