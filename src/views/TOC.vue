<template>
  <el-container class="home_container">
    <iframe src="/christmas.html" frameborder="0"
            style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; z-index: -1;"></iframe>
    <el-header class="home_header">
      <dev :class="{ hovered: drawer }"
           style="cursor: pointer;"
           @click="$router.push('/')"
           @mouseover="drawer=true"
           @mouseout="drawer=false">
        鸿雁长飞光不度
        <img src="@/assets/code.png" style="height: 0.7em" alt="NO">
        鱼龙浅跃水成文
      </dev>
    </el-header>
    <el-main class="home_main">Main</el-main>
    <el-footer class="home_footer">Copyright © 2023 Peter Cheung 保留所有权利</el-footer>
  </el-container>
</template>

<script>
export default {
  name: 'TOC',
  data() {
    return {
      drawer: false
    };
  },
  created() {
  },
  methods: {}
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
  background-color: transparent;
  color: white;
  text-align: center;
  align-items: center;
  justify-content: space-between;
}

.home_footer {
  background-color: transparent;
  text-align: center;
  color: rgb(142, 142, 145);
  font-size: 16px;
  line-height: 60px;
}
</style>