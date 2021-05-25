Squizit is a simple tool, that aim to help you get the grade you want, not the one you have learnt for.

[![Build Status](https://ci.mrcyjanek.net/badge/build-squizit.svg)](https://ci.mrcyjanek.net/jobs/build-squizit)

# Screenshots

| First, input PIN | Then enjoy! |
| ---------------- | ----------- |
| ![Fill me senpai](https://mrcyjanek.net/projects/squizit/static/input-pin.png) | ![Here it comes](https://mrcyjanek.net/projects/squizit/static/answers.png) |

# Hosted version

List of web instances [Host your own!](#host-on-your-site)

<!--
  Add new instances to the bottom of this list.
-->
 - [squizit.sivaj.pl](https://squizit.sivaj.pl/) <span id="squizitsivajpl"></span>
 - [squizit.cf](https://squizit.cf) <span id="squizitcf"></span>
 - Backend Server <span id="backend"></span>
 - Downloads Server <span id="downloads"></span>

Do you want your site to get added here? Submit a pull request to [git.mrcyjanek.net](https://git.mrcyjanek.net/mrcyjanek/mysite) or [github.com](https://github.com/MrCyjaneK/mysite)

<script>
function updateVersion(url, name) {
  let started = new Date().getTime();
  fetch(url+"/api/version")
    .then(r => r.text())
    .then(r => {
      let ended = new Date().getTime();
      let ms = Number(ended - started).toFixed(0);
      document.getElementById(name).innerHTML = "<span style=\"color: green;\">online <span style=\"color: yellow;\">["+ms+" ms]</span></span> "+r.substr(0,32).replaceAll(/&/g, "&amp;")
      .replaceAll(/</g, "&lt;")
      .replaceAll(/>/g, "&gt;")
      .replaceAll(/"/g, "&quot;")
      .replaceAll(/'/g, "&#039;")
    })
    .catch(e => {
      console.log(e.toString())
      let ended = new Date().getTime();
      let ms = Number(ended - started).toFixed(0);
      document.getElementById(name).innerHTML = "<span title=\""+e.toString().replaceAll('"', '')+"\" style=\"color: red;\">offline <span style=\"color: yellow;\">["+ms+" ms]</span> <b onclick=\"alert('"+e.toString().replaceAll('"', '')+"')\" style=\"color: white\">[?]</b></span>"
    })
}
setTimeout(() => updateVersion("https://squizit.sivaj.pl", "squizitsivajpl"))
setTimeout(() => updateVersion("https://squizit.cf", "squizitcf"))
//setTimeout(() => updateVersion("https://beta.squizit.cf", "betasquizitcf"))
setTimeout(() => updateVersion("https://squiz.mrcyjanek.net", "backend"))
setTimeout(() => updateVersion("https://static.mrcyjanek.net/laminarci/build-squizit/latest", "downloads"))
</script>

# Downloads

| ![Android](/static/icons/android-icon.svg) | ![Ubuntu Touch](/static/icons/ubuntu-icon.svg) | ![Micro$oft Windows](/static/icons/microsoft-icon.svg) | ![Debian/Ubuntu Package](/static/icons/debian-icon.svg) | ![MacOS Executable](/static/icons/apple-tile.svg) |
| --- | --- | --- | --- | --- |
| [.apk (all)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit.android.all.apk) | | | Instructions below. |
|  | [.click (armhf)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_arm.click) | | [Binary (armhf)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_armhf) |
|  | [.click (aarch64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_arm64.click) | | [Binary (aarch64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_arm64) | **unavailable** |
|  |  | [exe portable (x86)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_windows_386.exe) | [Binary (i386)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_386) |
|  | [.click (x86_64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_amd64.click) | [exe (x86_64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_windows_amd64.exe) | [Binary (amd64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_amd64) | [.zip (amd64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_darwin_amd64.zip) |

# Debian package

If you run on debian (or ubuntu/mint) machine, and would like to get automatic updates, you can install squizit directly from my apt repository.

First, install my repository

```plain
# wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
  apt install ./mrcyjanek-repo-latest.deb && \
  rm ./mrcyjanek-repo-latest.deb && \
  apt update
```

Then install squizit

```plain
# apt install squizit
```

# Host on your site

If you own a small server, you can help me with hosting the cheat! Simply run this command:

```plain
wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
  apt install ./mrcyjanek-repo-latest.deb && \
  rm ./mrcyjanek-repo-latest.deb && \
  apt update
apt install squizit squizit-server
```

And you will have a cheat running on your server!


\* note about android build.

Due to an upstream gradle issue ([1](https://github.com/gradle/gradle/issues/14528), [2](https://github.com/gradle/gradle/issues/12731)) I have to build on my pc, so updates may be pushed with little delay.
