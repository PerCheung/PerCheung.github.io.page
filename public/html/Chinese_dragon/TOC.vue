<template>
  <el-container class="home_container">
    <el-header class="home_header">
      <div class="header_left">
        <el-button
            icon="el-icon-back"
            size="small"
            @click="$router.push('/')"
            style="color:#bb8c2d">
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
    <el-main class="home_main"
             v-loading="loading"
             element-loading-text="正在查找文章"
             element-loading-background="transparent">
      <div class="home_main_toc">
        <div class="custom_card" v-for="item in filteredToc" :key="item">
          <el-link :href="getLink(item)"
                   :underline="false"
                   target="_blank"
                   style="color: #edd78f">
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
  color: #6e6e73;
  font-size: 0.8em;
  line-height: 80px;
}

@media screen and (max-width: 1023px) {
  .home_container {
    background-image: url("https://store.storeimages.cdn-apple.com/8756/as-images.apple.com/is/cny-header-202401-dragon-clouds?wid=3840&hei=400&fmt=png-alpha&.v=1704936825141");
    background-repeat: no-repeat;
    background-size: 300%;
    background-position: top;
    background-color: #f5f5f7;
    color: #cfa25c;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
  }

  .header_left {
    display: none;
  }

  .custom_card {
    width: 72%;
    margin: 0 auto 15px auto;
    padding: 13px 30px;
    background-color: rgba(158, 13, 29, 0.95);
    border-left: 5px solid #a5803a;
    border-radius: 4px;
    transition: box-shadow 0.3s ease-in-out;
  }

  .custom_card:hover {
    background-color: #9e0d1d;
    box-shadow: 1px 5px 10px rgba(0, 0, 0, 0.3);
  }

  .header_center {
    color: #cfa25c;
    width: 100%;
    margin: 10px auto;
    font-size: 1em;
    padding-bottom: 10px;
    transition: font-size 0.3s ease;
  }

  .header_center:hover {
    font-size: 1.3em;
  }

  .header_center:hover img {
    height: 0.84em;
  }

  .header_right {
    display: none;
  }

  .header_right_mobile {
    width: 80%;
    margin: 0 auto;
  }
}

@media screen and (min-width: 1024px) {
  .home_container {
    background-image: url("https://store.storeimages.cdn-apple.com/8756/as-images.apple.com/is/cny-header-202401-dragon-clouds?wid=3840&hei=400&fmt=png-alpha&.v=1704936825141");
    background-repeat: no-repeat;
    background-size: 140%;
    background-position: top;
    background-color: #f5f5f7;
    color: #cfa25c;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
  }

  .header_left {
    display: flex;
    align-items: center;
    width: 15%;
  }

  .custom_card {
    width: 58%;
    margin: 0 auto 20px auto;
    padding: 20px 30px;
    background-color: rgba(158, 13, 29, 0.95);
    border-left: 5px solid #a5803a;
    border-radius: 4px;
    transition: box-shadow 0.3s ease-in-out;
  }

  .custom_card:hover {
    background-color: #9e0d1d;
    box-shadow: 1px 5px 10px rgba(0, 0, 0, 0.3);
  }

  .header_center {
    color: #cfa25c;
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