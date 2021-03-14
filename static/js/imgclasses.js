new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue.js!',
      showImgUrl: "https://dss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1473836766,4030812874&fm=26&gp=0.jpg",
      items:[
        {
            id: "",
            classes: "未分类",
            url: "https://dss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1473836766,4030812874&fm=26&gp=0.jpg",
        },
        {
            id: "",
            classes: "未分类",
            url: "https://scpic.chinaz.net/Files/pic/pic9/202103/apic31315_s.jpg",
        },
      ],
      buttons:[
        "正常",
        "异常",
      ]
    },
    methods: {
        changeitem(it){
            this.showImgUrl=it.url;
        }
    }
})