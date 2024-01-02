<template>
  <el-container class="home_container">
    <iframe src="/html/christmas.html" frameborder="0"
            style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; z-index: -1;"></iframe>
    <el-header class="home_header">
      <div class="header_left">
        <el-button
            icon="el-icon-back"
            size="small"
            @click="$router.push('/')"
            type="primary" plain>
          返回我的主页
        </el-button>
      </div>
      <dev class="header_center"
           @mouseover="changeFontSize(true)"
           @mouseout="changeFontSize(false)">
        <img src="@/assets/toc.png" style="height: 0.7em" alt="@TOC">&nbsp;
        我的博客
      </dev>
      <div class="header_right">
        <el-input
            size="small"
            v-model="search"
            prefix-icon="el-icon-search"
            placeholder="搜索文章">
        </el-input>
      </div>
    </el-header>
    <hr style="
              border: none;
              height: 1px;
              width: 58%;
              background-color: #eaecef;
              margin: 0 auto;
    ">
    <el-main class="home_main">
      <div class="home_main_toc">
        <div class="custom_card" v-for="item in filteredToc" :key="item">
          <el-link :href="getLink(item)"
                   :underline="false"
                   target="_blank">
            <i class="el-icon-paperclip"></i>
            {{ item }}
          </el-link>
        </div>
      </div>
    </el-main>
    <el-footer class="home_footer">Copyright © 2023 Peter Cheung 保留所有权利</el-footer>
  </el-container>
</template>

<script>
export default {
  name: 'TOC',
  data() {
    return {
      isHovered: false,
      search: '',
      toc: ["HTML5学习记录", "markdown各种语法测试", "初步领悟C指针", "卓别林的演讲", "如何搭建github.io网页", "搭建多人聊天室（保姆级教学，从0开始）", "正则表达式笔记"]
    };
  },
  computed: {
    filteredToc() {
      return this.toc.filter(item =>
          !this.search || item.toLowerCase().includes(this.search.toLowerCase())
      );
    }
  },
  methods: {
    changeFontSize(isHovered) {
      this.isHovered = isHovered;
    },
    getLink(item) {
      return `https://percheung.github.io/blog/${item}`;
    }
  }
}
</script>

<style scoped>
.home_container {
  background-color: #acbbcc;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  color: black;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
}

.home_header {
  background-color: transparent;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 10px 20px;
}

.header_left {
  display: flex;
  align-items: center;
  width: 15%;
}

.header_center {
  color: #2768d7;
  font-weight: bold;
  display: flex;
  align-items: center;
  font-size: 1.2em;
  transition: font-size 0.3s ease;
}

.header_center:hover {
  font-size: 1.5em;
}

.header_center:hover img {
  height: 0.84em;
}

.header_right {
  display: flex;
  align-items: center;
  width: 15%;
}

.home_footer {
  background-color: transparent;
  text-align: center;
  color: rgb(142, 142, 145);
  font-size: 16px;
  line-height: 60px;
}

.home_main_toc {
  width: 60%;
  margin: 0 auto;
}

.custom_card {
  margin-top: 0;
  margin-bottom: 20px;
  padding: 20px 30px;
  background-color: #edf8ff;
  border-left: 5px solid #63c0ff;
  border-radius: 4px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.2s ease-in-out;
}

.custom_card:hover {
  box-shadow: 0 8px 8px rgba(0, 0, 0, 0.2);
}
</style>