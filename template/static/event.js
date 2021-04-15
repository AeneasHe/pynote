String.prototype.endWith = function (s) {
    var d = this.length - s.length;
    return d >= 0 && this.lastIndexOf(s) == d;
};

const counter = document.querySelector(".counter");
const pathList = document.querySelector(".path-container");

const btnIncr = document.querySelector(".btn-incr");
const btnDecr = document.querySelector(".btn-decr");

// 文件和路径相关按钮
const btnFolder = document.querySelector(".btn-folder");
const btnFile = document.querySelector(".btn-file");
const btnAll = document.querySelector(".btn-all");

const btnOpenWindow = document.querySelector(".btn-openwindow");

const fileContent = document.querySelector(".file-content");

// const render = async () => {
//   counter.innerText = `Count: ${await window.counterValue()}`;
// };

const toDom = (str) => {
    var dom,
        tmp = document.createElement("div");
    tmp.innerHTML = str;
    dom = tmp.children[0];
    return dom;
};

const appendToPathList = (str) => {
    console.log("===>4", str)
    var name = str.replace(/^.*[\\\/]/, "");
    var filename = name.replace(".md", "");
    var pathType = "path-file";

    // if (name.endWith(".md")) {
    //   pathType = "path-folder";
    // }

    if (name == filename) {
        pathType = "path-folder";
    }

    pathList.appendChild(
        toDom(
            '<div class="path-name-row ' +
            pathType +
            '"><div class="path-name" title="' +
            str +
            '" data-type="' + pathType + '"">' +
            filename +

            "</div></div> "
        )
    );
};


btnFile.addEventListener("click", async () => {
    var file = await currentFile();
    fileContent.innerText = file;
});

//   btnIncr.addEventListener("click", async () => {
//     await counterAdd(1); 
//     render();
//   });

//   btnDecr.addEventListener("click", async () => {
//     await counterAdd(-1);
//     render();
//   });

btnFolder.addEventListener("click", async () => {
    folders = await showPath(".", "folder");
    pathList.innerText = "";
    folders.map((f) => appendToPathList(f));
});

btnFile.addEventListener("click", async () => {
    folders = await showPath(".", "file");
    pathList.innerText = "";
    folders.map((f) => appendToPathList(f));
});

btnAll.addEventListener("click", async () => {
    folders = await showPath(".", "all");
    pathList.innerText = "";
    folders.map((f) => appendToPathList(f));
});

// btnOpenWindow.addEventListener("click", async () => {
//   await openWindow(""); 
// });

// 默认显示文件路径
showPath(".", "file").then((files) => {
    pathList.innerText = "";
    files.map((f) => appendToPathList(f));
});

// 获取并显示当前文档内容
currentFile().then((file) => {
    fileContent.innerText = file;
});

// 点击路径时处理
$(document).on("click", ".path-name", async (event) => {
    var target = $(event.target);

    // 其他恢复默认颜色
    $(".path-name-row").css("background", "white");
    // 点击的设置为红色
    target.parent().css("background", "red");

    var pathType = target.attr("data-type")
    var pathname = target.attr("title");

    // webview注入的函数

    console.log("===>1");
    if (pathType == "path-file") {
        await openPath(pathname, pathType);
        currentFile().then((file) => {
            fileContent.innerText = file;
        });
    }

    console.log("===>2");
    if (pathType == "path-folder") {
        folders = await showPath(pathname, "all");

        pathList.innerText = "";

        console.log("===>3", folders);

        folders.map((f) => appendToPathList(f));
    }


});
    //render();
