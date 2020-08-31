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
        <el-form-item label="佣金比(小于)">
          <el-input type='number' placeholder="佣金比" v-model.number="searchInfo.fee_rate"></el-input>
        </el-form-item>
        <el-form-item label="品类">
          <el-select clearable placeholder="请选择" v-model="searchInfo.third_category">
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
        <el-form-item>
          <el-button @click="openDialog('add')" type="primary">新增商品</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="id" min-width="50" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="sku" min-width="80" sortable="custom">
        <template slot-scope="scope">
          <a :href="'http://item.jd.com/' + scope.row.sku_id + '.html'" target="_blank" style="color:#606266">
            {{scope.row.sku_id}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="商品名" min-width="200" sortable="custom">
        <template slot-scope="scope">
          <a :href="'https://union.jd.com/proManager/index?pageNo=1&keywords=' + scope.row.sku_id" target="_blank" style="color:#606266">
            {{scope.row.sku_name}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="价格" min-width="50" prop="price" sortable="custom"></el-table-column>
      <el-table-column label="佣金比" min-width="50" prop="fee_rate" sortable="custom"></el-table-column>
      <el-table-column label="佣金" min-width="50" prop="fee" sortable="custom"></el-table-column>
      <el-table-column label="品类" min-width="50" prop="third_category" sortable="custom"></el-table-column>
      <el-table-column label="是否自营" min-width="50" prop="jd_sale" sortable="custom"></el-table-column>

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editGoods(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteGoods(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
        </template>
      </el-table-column>
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

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="goodsForm">
        <el-form-item label="sku" prop="sku_id">
          <el-input autocomplete="off" v-model="form.sku_id"></el-input>
        </el-form-item>
        <el-form-item label="商品名" prop="sku_name">
          <el-input autocomplete="off" v-model="form.sku_name"></el-input>
        </el-form-item>
      </el-form>
      <div class="edit_container">
        <quill-editor
                :options="{}"
                ref="myQuillEditor"
                v-model="form.content"
        ></quill-editor>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getGoodsList,
  getGoodsById,
  updateGoods,
  createGoods,
  deleteGoods
} from '@/api/goods'

import imageResize from 'quill-image-resize-module'
Quill.register('modules/imageResize', imageResize)

import infoList from '@/components/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
const methodOptions = [
  {
    value: '乳胶枕',
    label: '乳胶枕',
    type: ''
  },
  {
    value: '按摩器',
    label: '按摩器',
    type: ''
  },
  {
    value: '按摩椅',
    label: '按摩椅',
    type: ''
  }
]

export default {
  name: 'Goods',
  mixins: [infoList],
  data() {
    return {
      listApi: getGoodsList,
      dialogFormVisible: false,
      dialogTitle: '新增商品',
      form: {
        sku_id: '',
        sku_name: '',
        content: ''
      },
      methodOptions: methodOptions,
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
    },
    initForm() {
      this.$refs.goodsForm.resetFields()
      this.form= {
        sku_id: '',
        sku_name: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'add':
          this.dialogTitle = '新增商品'
          break
        case 'edit':
          this.dialogTitle = '编辑商品'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async editGoods(row) {
      const res = await getGoodsById({ id: row.ID })
      this.form = res.data.goods
      this.openDialog('edit')
    },
    async deleteGoods(row) {
      this.$confirm('此操作将删除此商品, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteGoods(row)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    async enterDialog() {
      this.$refs.goodsForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              {
                const res = await createGoods(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }

              break
            case 'edit':
              {
                const res = await updateGoods(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
              break
            default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
              break
          }
        }
      })
    }
  },
  filters: {
    methodFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`
    },
    tagTypeFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    }
  },
  created(){
    this.getTableData()
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