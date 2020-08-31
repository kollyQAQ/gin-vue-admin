<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="问题ID">
          <el-input placeholder="问题ID" v-model="searchInfo.qid"></el-input>
        </el-form-item>
        <el-form-item label="问题名称">
          <el-input placeholder="问题名称" v-model="searchInfo.qname"></el-input>
        </el-form-item>
        <el-form-item label="品类">
          <el-select clearable placeholder="请选择" v-model="searchInfo.type">
            <el-option
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
              v-for="item in typesOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('add')" type="primary">新增问答</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="ID" min-width="50" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="问题ID" min-width="80" sortable="custom">
        <template slot-scope="scope">
          <a :href="'https://www.zhihu.com/question/' + scope.row.qid" target="_blank" style="color:#606266">
            {{scope.row.qid}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="问题名称" min-width="180" prop="qname" sortable="custom"></el-table-column>
      <el-table-column label="总浏览" min-width="80" prop="view_total" sortable="custom"></el-table-column>
      <el-table-column label="总回答" min-width="60" prop="answer_total" sortable="custom"></el-table-column>
      <el-table-column label="今日浏览" min-width="80" prop="today_add_view" sortable="custom"></el-table-column>
      <el-table-column label="三日浏览" min-width="80" prop="three_day_add_view" sortable="custom"></el-table-column>
      <el-table-column label="回答ID" min-width="80" sortable="custom">
        <template slot-scope="scope">
          <a :href="'https://www.zhihu.com/question/' + scope.row.qid + '/answer/' + scope.row.aid" target="_blank" style="color:#606266">
            {{scope.row.aid}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="排名" min-width="60" prop="rank" sortable="custom"></el-table-column>
      <el-table-column label="点赞" min-width="60" prop="like_num" sortable="custom"></el-table-column>
      <el-table-column label="带货" min-width="60" prop="with_card" sortable="custom"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editQa(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="viewHistory(scope.row)" size="small" type="primary" icon="el-icon-edit">历史</el-button>
          <el-button @click="deleteQa(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
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
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="问题ID" prop="qid">
          <el-input autocomplete="off" v-model="form.qid"></el-input>
        </el-form-item>
        <el-form-item label="问题名称" prop="qname">
          <el-input autocomplete="off" v-model="form.qname"></el-input>
        </el-form-item>
        <el-form-item label="回答ID" prop="aid">
          <el-input autocomplete="off" v-model="form.aid"></el-input>
        </el-form-item>
        <el-form-item label="是否带货" prop="with_card">
          <el-select placeholder="请选择" v-model="form.with_card">
            <el-option
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :before-close="closeHisDialog" :title="dialogTitle" :visible.sync="dialogHisVisible">
      <div>
        <div id="main" class="chart-container"></div>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeHisDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getQaList,
  getQaById,
  createQa,
  updateQa,
  deleteQa,
  queryQaHistory,
} from '@/api/qa'

import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import infoList from '@/components/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'

const typesOptions = [
  {
    value: 1,
    label: '礼物',
    type: ''
  },
  {
    value: 2,
    label: '乳胶枕',
    type: ''
  },
  {
    value: 3,
    label: '按摩仪',
    type: ''
  }
]

const methodOptions = [
  {
    value: '0',
    label: '不带货',
    type: ''
  },
  {
    value: '1',
    label: '带货',
    type: ''
  }
]

export default {
  name: 'Qa',
  mixins: [infoList],
  data() {
    return {
      listApi: getQaList,
      dialogFormVisible: false,
      dialogHisVisible: false,
      dialogTitle: '新增问答',
      form: {
        qid: '',
        aid: '',
        with_card: '0'
      },
      methodOptions: methodOptions,
      typesOptions: typesOptions,
      type: '',
      rules: {
        qid: [{ required: true, message: '请输入问题ID', trigger: 'blur' }],
        aid: [
          { required: false, message: '请输入回答ID', trigger: 'blur' }
        ],
        with_card: [
          { required: false, message: '回答是否带货', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order === 'descending'
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
      this.$refs.apiForm.resetFields()
      this.form= {
        qid: '',
        qname:'',
        type:0
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    closeHisDialog() {
      this.dialogHisVisible = false
    },
    openDialog(type) {
      console.log(type)
      switch (type) {
        case 'add':
          this.dialogTitle = '新增问答'
          this.dialogFormVisible = true
          break
        case 'edit':
          this.dialogTitle = '编辑问答'
          this.dialogFormVisible = true
          break
        case 'history':
          this.dialogTitle = '历史趋势'
          this.dialogHisVisible = true
          break
        default:
          break
      }
      this.type = type
    },
    async viewHistory(row) {


      this.dialogHisVisible = true


      let myChart = echarts.init(document.getElementById('main'));

      let  option = {
        legend: {},
        tooltip: {
          trigger: 'axis',
          showContent: false
        },
        dataset: {
          source: [
            ['product', '2012', '2013', '2014', '2015', '2016', '2017','2018','2019','2020'],
            ['Matcha Latte', 41.1, 30.4, 65.1, 53.3, 83.8, 70.0,6.4, 65.2, 82.5],
            ['Milk Tea', 86.5, 92.1, 85.7, 83.1, 73.4, 55.1,2, 67.1, 69.2],
            ['Cheese Cocoa', 24.1, 67.2, 79.5, 86.4, 65.2, 82.5,65.1, 53.3, 83.8],
            ['Walnut Brownie', 55.2, 67.1, 69.2, 72.4, 53.9, 39.1,86.5, 92.1, 85.7]
          ]
        },
        xAxis: {
          type: 'category',
          axisLabel: {
            show: true,
            textStyle: {
              color: 'rgb(192,192,192)',  //更改坐标轴文字颜色
              fontSize : 14    //更改坐标轴文字大小
            }
          },
          axisTick: {
            show: false
          },
          axisLine:{
            lineStyle:{
              color:'rgb(192,192,192)' //更改坐标轴颜色
            }
          },
        },
        yAxis: {
          gridIndex:0,
          axisLabel: {
            show: true,
            textStyle: {
              color: 'rgb(192,192,192)',  //更改坐标轴文字颜色
              fontSize: 14    //更改坐标轴文字大小
            }
          },
          axisTick: {
            show: false
          },
          axisLine: {
            lineStyle: {
              color: 'rgb(192,192,192)' //更改坐标轴颜色
            }
          }
        },
        grid: {top: '55%'},
        series: [
          {type: 'line', smooth: true, seriesLayoutBy: 'row'},
          {type: 'line', smooth: true, seriesLayoutBy: 'row'},
          {type: 'line', smooth: true, seriesLayoutBy: 'row'},
          {type: 'line', smooth: true, seriesLayoutBy: 'row'},
          {
            type: 'pie',
            id: 'pie',
            radius: '30%',
            center: ['50%', '25%'],
            label: {
              formatter: '{b}: {@2012} ({d}%)'
            },
            encode: {
              itemName: 'product',
              value: '2012',
              tooltip: '2012'
            }
          }
        ]
      };

      // window.addEventListener('resize',function() {myChart.resize()});
      myChart.setOption(option);

      const res = await queryQaHistory({ id: row.qid })
      this.form = res.data.qa
      this.openDialog('history')
    },
    async editQa(row) {
      const res = await getQaById({ id: row.qid })
      this.form = res.data.qa
      this.openDialog('edit')
    },
    async deleteQa(row) {
      this.$confirm('此操作将删除此问答, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteQa(row)
          if (res.code === 0) {
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
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              {
                const res = await createQa(this.form)
                if (res.code === 0) {
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
                const res = await updateQa(this.form)
                if (res.code === 0) {
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
            case 'history':
              {
                const res = await updateQa(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.closeHisDialog()
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