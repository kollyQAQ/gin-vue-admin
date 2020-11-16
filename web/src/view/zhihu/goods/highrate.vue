<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="sku">
          <el-input placeholder="sku" v-model="searchInfo.sku_id"></el-input>
        </el-form-item>
        <el-form-item label="商品名">
          <el-input placeholder="商品名" v-model="searchInfo.sku_name"></el-input>
        </el-form-item>
<!--        <el-form-item label="佣金比(小于)">-->
<!--          <el-input type='number' placeholder="佣金比" v-model.number="searchInfo.fee_rate"></el-input>-->
<!--        </el-form-item>-->
        <el-form-item label="一级类目">
          <el-select clearable placeholder="请选择" v-model="searchInfo.category_one">
            <el-option
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="二级类目">
          <el-select clearable placeholder="请选择" v-model="searchInfo.category_two">
            <el-option
                    :key="item.value"
                    :label="`${item.label}`"
                    :value="item.value"
                    v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="三级类目">
          <el-select clearable placeholder="请选择" v-model="searchInfo.category_three">
            <el-option
                    :key="item.value"
                    :label="`${item.label}`"
                    :value="item.value"
                    v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="id" min-width="50" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="sku" min-width="80" prop="sku_id" sortable="custom">
        <template slot-scope="scope">
          <a :href="'http://item.jd.com/' + scope.row.sku_id + '.html'" target="_blank" style="color:#606266">
            {{scope.row.sku_id}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="商品名" min-width="200" prop="goods_name" sortable="custom">
        <template slot-scope="scope">
          <a :href="scope.row.url" target="_blank" style="color:#606266">
            {{scope.row.goods_name}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="价格" min-width="50" prop="price" sortable="custom"></el-table-column>
      <el-table-column label="佣金比" min-width="50" prop="fee_rate" sortable="custom">
        <template slot-scope="scope">
          <p v-if="scope.row.fee_rate>=20" style="font-size:16px;color:red">{{scope.row.fee_rate}}</p>
          <p v-else-if="scope.row.fee_rate>=10" style="font-size:16px;color:orange">{{scope.row.fee_rate}}</p>
          <p v-else>{{scope.row.fee_rate}}</p>
        </template>
      </el-table-column>
      <el-table-column label="佣金" min-width="50" prop="fee" sortable="custom"></el-table-column>
      <el-table-column label="一级类目" min-width="50" prop="category_one" sortable="custom"></el-table-column>
      <el-table-column label="二级类目" min-width="50" prop="category_two" sortable="custom"></el-table-column>
      <el-table-column label="三级类目" min-width="50" prop="category_three" sortable="custom"></el-table-column>
      <el-table-column label="是否自营" min-width="50" prop="jd_sale" sortable="custom">
        <template slot-scope="scope">
          <p v-if="scope.row.jd_sale == 1">
            <i class="el-icon-success" style="font-size:20px;color: green"></i>
          </p>
          <p v-if="scope.row.jd_sale == 0"><i class="el-icon-error"></i></p>
        </template>
      </el-table-column>
      <el-table-column label="最后更新时间" min-width="100" prop="update_time"></el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[20, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getHighRateGoodsList,
  getThirdCategory
} from '@/api/goods'


import infoList from '@/components/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'

export default {
  name: 'Goods',
  mixins: [infoList],
  data() {
    return {
      listApi: getHighRateGoodsList,
      methodOptions: [],
      type: '',
      rules: {
        sku_id: [{ required: true, message: '请输入sku_id', trigger: 'blur' }],
        sku_name: [
          { required: false, message: '请输入商品名称', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order == 'descending'
      }
      this.getTableData()
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 20
      this.getTableData()
    }
  },
  async created(){
    this.getTableData()
    const res = await getThirdCategory()
    this.methodOptions = res.data
  }
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>
