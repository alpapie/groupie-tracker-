mobilebtn=document.getElementById("onpen-menu")
mobileMenu=document.getElementById("onclick-mobile-btn")
closebtn=document.getElementById("closebtn")

header=document.getElementById("after-loading")
mainpage=document.getElementById("after-loading1")
loader=document.getElementById("loader")


setTimeout(()=>{
    loader.style.display="none"
    loader.style.display="hidden"  

    mainpage.style.display="block"
    mainpage.style.visibility="visible"

    header.style.display="block"
    header.style.visibility="visible"
},4000)

mobilebtn.addEventListener("click",()=>{
    mobileMenu.style.display="block"
    mobileMenu.style.visibility="visible"
    mobileMenu.style.marginLeft=0
    mobilebtn.style.display="none"
    mobilebtn.style.display="hidden"
},false)

closebtn.addEventListener("click",()=>{
    mobilebtn.style.display="block"
    mobilebtn.style.visibility="visible"
    // mobileMenu.style.display="none"
    // mobileMenu.style.display="hidden"
    mobileMenu.style.marginLeft="-100%"
})
