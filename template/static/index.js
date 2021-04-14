function setText(id, t) {
  document.getElementById(id).innerText = '页面将在' + t + '秒后自动跳转';
}

function append(str){
    var span = document.createElement('span')
    span.innerText = str
    document.body.appendChild(span)
}

function time(id, cb) {
  var t = 5;
  var timer = setInterval(function () {
    t -= 1;
    setText(id, t);
    if (t <= 0) {
      clearInterval(timer);
      if (cb) cb();
    }
  }, 1000);
}

function goTo(url) {
  window.location.href = url;
  // console.log(url);
}