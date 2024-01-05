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
    <div class="header_right_mobile">
      <el-input
          v-model="search"
          size="small"
          prefix-icon="el-icon-search"
          placeholder="搜索文章">
      </el-input>
    </div>
    <hr class="header_hr">
    <el-main class="home_main"
             v-loading="loading"
             element-loading-text="正在查找文章"
             element-loading-background="transparent">
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
      toc: [],
      loading: true
    };
  },
  computed: {
    filteredToc() {
      return this.toc.filter(item =>
          !this.search || item.toLowerCase().includes(this.search.toLowerCase())
      );
    }
  },
  created() {
    this.fetchTocData();
  },
  methods: {
    changeFontSize(isHovered) {
      this.isHovered = isHovered;
    },
    getLink(item) {
      return `https://percheung.github.io/blog/${item}`;
    },
    async fetchTocData() {
      try {
        const response = await fetch('https://percheung.github.io/blog/toc.json');
        const jsonData = await response.json();

        // 对jsonData进行排序
        jsonData.sort((a, b) => {
          if (a[0] !== b[0]) {
            // 简体中文的排序
            return a.localeCompare(b, 'zh-Hans-CN', {sensitivity: 'accent'});
          } else {
            // 首字母相等时按照长度排序
            return a.length - b.length;
          }
        });

        this.toc = jsonData;

        this.loading = false;
      } catch (error) {
        console.error('Error fetching toc data:', error);
      }
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

.home_footer {
  background-color: transparent;
  text-align: center;
  color: rgb(142, 142, 145);
  font-size: 0.8em;
  line-height: 80px;
}

@media screen and (max-width: 1023px) {
  .header_left {
    display: none;
  }

  .header_hr {
    border: none;
    height: 1px;
    width: 80%;
    background-color: #eaecef;
    margin: 0 auto;
  }

  .custom_card {
    width: 72%;
    margin: 0 auto 15px auto;
    padding: 13px 30px;
    background-color: #edf8ff;
    border-left: 5px solid #63c0ff;
    border-radius: 4px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.2s ease-in-out;
  }

  .custom_card:hover {
    box-shadow: 0 8px 8px rgba(0, 0, 0, 0.2);
  }

  .header_center {
    color: #2768d7;
    width: 100%;
    margin: 10px auto;
    font-weight: bold;
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
    display: none;
  }

  .header_right_mobile {
    width: 80%;
    margin: -5px auto 10px auto;
  }
}

@media screen and (min-width: 1024px) {
  .header_left {
    display: flex;
    align-items: center;
    width: 15%;
  }

  .custom_card {
    width: 58%;
    margin: 0 auto 20px auto;
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

  .header_hr {
    border: none;
    height: 1px;
    width: 60%;
    background-color: #eaecef;
    margin: 0 auto;
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

  .header_right_mobile {
    display: none;
  }
}
</style>