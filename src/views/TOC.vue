<template>
  <el-container class="home_container">
    <iframe src="/christmas.html" frameborder="0"
            style="position: fixed; top: 0; left: 0; width: 100%; height: 100%; z-index: -1;"></iframe>
    <el-header class="home_header">
      <dev :class="{ hovered: true }"
           style="cursor: pointer;"
           @click="$router.push('/')">
        鸿雁长飞光不度
        <img src="@/assets/code.png" style="height: 0.7em" alt="NO">
        鱼龙浅跃水成文
      </dev>
    </el-header>
    <el-main class="home_main">
      <div class="table-container">
        <el-table
            :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
            style="width:100%;">
          <el-table-column prop="name">
            <template slot="header" slot-scope="scope">
              <div class="centered-content">
                <span :class="{ 'larger-font': isHovered }"
                      @mouseover="isHovered=true"
                      @mouseout="isHovered=false">
                  <img src="@/assets/toc.png" style="height: 0.8em" alt="@TOC">
                  我的博客文章目录
                </span>
              </div>
            </template>
          </el-table-column>
          <el-table-column
              width="150px"
              align="center">
            <template slot="header" slot-scope="scope">
              <el-input
                  v-model="search"
                  size="small"
                  prefix-icon="el-icon-search"
                  placeholder="搜索文章"/>
            </template>
            <template slot-scope="scope">
              <!--<el-button-->
              <!--  size="small"-->
              <!--  type="primary"-->
              <!--  @click="clickmd(scope.$index, scope.row)">阅读-->
              <!--</el-button>-->
              <el-button
                  size="small"
                  type="primary">阅读
              </el-button>
            </template>
          </el-table-column>
        </el-table>
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
      tableData: [{
        name: '李小虎'
      }, {
        name: '王小虎'
      }, {
        name: '马小虎'
      }, {
        name: '王小虎'
      }],
      search: ''
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

.table-container {
  width: 80%;
  margin: 0 auto;
}

.centered-content {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 10px 0;
  font-size: 1.5em;
}

.larger-font {
  font-size: 1.2em;
}

/deep/ .table-container .el-table,
/deep/ .table-container .el-table__expanded-cell {
  background-color: transparent;
}

/deep/ .table-container .el-table th,
/deep/ .table-container .el-table tr,
/deep/ .table-container .el-table__header-wrapper th {
  background-color: transparent;
  color: white;
}

/deep/ .table-container .el-table td {
  background-color: transparent;
  border: none;
  font-weight: bold;
}

/deep/ .table-container .el-table .cell {
  line-height: 1.5em;
  font-size: 1.5em;
}

/deep/ .table-container .el-table::before {
  display: none;
}
</style>