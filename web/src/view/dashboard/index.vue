<template>
 <div class="big">
     <div>
         <el-form :model="goodsData">
             <el-form-item label="dddd">
                 <el-input placeholder="搜索条件" v-model="goodsData.aaa"></el-input>
             </el-form-item>
         </el-form>
     </div>
     <el-table :data="qaStatData" border stripe>
         <el-table-column label="问题类型" min-width="100" prop="type_desc"></el-table-column>
         <el-table-column label="问题总数" min-width="100" prop="total"></el-table-column>
         <el-table-column label="已回答问题数" min-width="100" prop="answer"></el-table-column>
         <el-table-column label="带货回答数" min-width="100" prop="card_answer"></el-table-column>
         <el-table-column label="回答占比" min-width="100">
             <template slot-scope="scope">
                 {{Math.ceil(scope.row.answer / scope.row.total * 100)}} %
             </template>
         </el-table-column>
         <el-table-column label="带货占比" min-width="100">
             <template slot-scope="scope">
                 {{Math.ceil(scope.row.card_answer / scope.row.answer * 100)}} %
             </template>
         </el-table-column>
     </el-table>
     <div class="mid">
         <el-row :gutter="32">
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <raddar-chart />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <stackMap />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <Sunburst/>
                 </div>
             </el-col>
         </el-row>
     </div>
     <div class="top">
         <div id="main" class="chart-container"></div>
     </div>
      <div class="bottom">
         <el-row :gutter="32">
             <el-col :xs="24" :sm="24" :lg="12">
                 <div class="chart-player">
                     <musicPlayer />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="12">
                 <div  class="chart-player">
                     <todo-list />
                 </div>
             </el-col>
         </el-row>
     </div>
 </div>

</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import RaddarChart from "./component/RaddarChart"
import stackMap from "./component/stackMap"
import Sunburst from "./component/Sunburst"
import musicPlayer from "./component/musicPlayer"
import TodoList from "./component/todoList"

import {
    queryQaStat
} from '@/api/stat'

export default {
  name: 'Dashboard',
  data() {
    return {
        qaStatData: [],
        goodsData: {
            aaa:1
        },
    }
  },
    components:{
        RaddarChart, //雷达图
        stackMap, //堆叠图
        Sunburst, //旭日图
        musicPlayer,  //音乐播放器
        TodoList //TodoList
    },
    async created(){
        const res = await queryQaStat();
        this.qaStatData = res.data;
    },
  mounted() {
      let myChart = echarts.init(document.getElementById('main'),'macarons');
      // let stackMap = echarts.init(document.getElementById('stackMap'));
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
      //点记标记点时的动效
      myChart.on('updateAxisPointer', function (event) {
          var xAxisInfo = event.axesInfo[0];
          if (xAxisInfo) {
              var dimension = xAxisInfo.value + 1;
              myChart.setOption({
                  series: {
                      id: 'pie',
                      label: {
                          formatter: '{b}: {@[' + dimension + ']} ({d}%)'
                      },
                      encode: {
                          value: dimension,
                          tooltip: dimension
                      }
                  }
              });
          }
      });

      window.addEventListener('resize',function() {myChart.resize()});
      myChart.setOption(option);

  }
}
</script>

<style lang="scss" scoped>
    .big{
        margin:100px 0 0 0;
        padding-top: 0;
        background-color: rgb(243,243,243);;
        .top{
            width: 100%;
            height: 360px;
            margin-top: 20px;
            overflow: hidden;
            .chart-container{
                position: relative;
                width: 100%;
                height: 100%;
                padding: 20px;
                background-color: #fff;
            }
        }
        .mid{
            width: 100%;
            height: 380px;
            .chart-wrapper {
                height: 340px;
                background: #fff;
                padding: 16px 16px 0;
                margin-bottom: 32px;
            }
        }
        .bottom{
            width: 100%;
            height: 300px;
            margin: 20px 0;
            .el-row{
                margin-right: 4px !important;
            }
            .chart-player{
                width: 100%;
                height: 270px;
                padding: 10px;
                background-color: #fff;
            }
        }
    }

</style>
