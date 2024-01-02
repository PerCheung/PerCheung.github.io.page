<template>
  <el-container class="home_container">
    <iframe src="/html/christmas.html" frameborder="0"
            style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; z-index: -1;"></iframe>
    <el-header class="home_header">
      <div class="header_left">
        <el-button
            style="color:black"
            icon="el-icon-s-home"
            size="small"
            @click="$router.push('/')"
            plain>
          返回我的主页
        </el-button>
      </div>
      <dev class="header_center">
        <img src="@/assets/toc.png" style="height: 0.7em" alt="@TOC">&nbsp;
        我的博客文章目录
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
    <el-main class="home_main">
      <div class="home_main_toc">
        <el-card class="custom_card" v-for="item in filteredToc" :key="item">
          <el-card class="custom_card1">
            <el-link style="color: white;font-size: 1.2em;font-weight: bold;"
                     :href="getLink(item)"
                     :underline="false"
                     target="_blank">
              {{ item }}
            </el-link>
          </el-card>
        </el-card>
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
      search: '',
      toc: [
        "初步领悟C指针",
        "搭建多人聊天室（保姆级教学，从0开始）",
        "如何搭建github.io网页"
      ]
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
  color: white;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
}

.home_header {
  background-color: #78334f;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px;
}

.header_left {
  display: flex;
  align-items: center;
  width: 15%;
}

.header_center {
  display: flex;
  align-items: center;
  font-size: 1.2em;
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
  background-color: #5f7f99;
  border: none;
  border-radius: 0;
  box-shadow: 0 5px 4px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.2s ease-in-out;
}

.custom_card:hover {
  box-shadow: 0 10px 8px rgba(0, 0, 0, 0.2);
}

.custom_card1 {
  background-color: #5f7f99;
  border: 2px solid #84b1d3;
  margin: -16px;
  border-radius: 0;
}
</style>