<template>
  <el-row :gutter="40" class="panel-group" style="margin-top: 0">
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel">
        <div class="card-panel-icon-wrapper icon-people">
          <svg-icon icon-class="peoples" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            今日浏览
          </div>
          <count-to :start-val="0" :end-val="stat.totalView" :duration="duration" class="card-panel-num" />
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel">
        <div class="card-panel-icon-wrapper icon-money">
          <svg-icon icon-class="money" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            预估收入
          </div>
          <count-to :start-val="0" :end-val="stat.totalFee" :duration="duration" :prefix='prefix' class="card-panel-num" />
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel">
        <div class="card-panel-icon-wrapper icon-shopping">
          <svg-icon icon-class="shopping" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            无佣商品
          </div>
          <div v-if="stat.noFeeGoodsNum < 1">
            <count-to :start-val="0" :end-val="stat.noFeeGoodsNum" :duration="duration" class="card-panel-num" />
          </div>
          <div v-else>
            <count-to :start-val="0" :end-val="stat.noFeeGoodsNum" :duration="duration" class="card-panel-num" style="color:red" />
          </div>
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel">
        <div class="card-panel-icon-wrapper icon-message">
          <svg-icon icon-class="skill" class-name="card-panel-icon" />
<!--          <svg-icon icon-class="message" class-name="card-panel-icon" />-->
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            待办任务
          </div>
          <count-to :start-val="0" :end-val="stat.todoNum" :duration="duration" class="card-panel-num" />
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import CountTo from 'vue-count-to'

import {
  queryStat
} from '@/api/stat'

export default {
  components: {
    CountTo
  },
  data() {
    return {
        prefix:'￥ ',
        duration: 1000,
        stat: {
          totalView:0,
          totalFee:0,
          noFeeGoodsNum:0,
          todoNum:0
        },
    };
  },
  methods: {
    handleSetLineChartData(type) {
      this.$emit('handleSetLineChartData', type)
    }
  },
  mounted(){
    if(this.timer){
      clearInterval(this.timer)
    }else{
      this.timer = setInterval(async ()=>{
        // 调用相应的接口，渲染数据
        const res = await queryStat();
        this.stat.totalView = res.data.stat.today_view;
        this.stat.totalFee = res.data.stat.today_fee;
        this.stat.noFeeGoodsNum = res.data.stat.no_fee_goods_num;
        this.stat.todoNum = res.data.stat.todo_number;
      },60000)
    }
  },
  destroyed(){
    clearInterval(this.timer)
  },
  async created(){
    const res = await queryStat();
    this.stat.totalView = res.data.stat.today_view;
    this.stat.totalFee = res.data.stat.today_fee;
    this.stat.noFeeGoodsNum = res.data.stat.no_fee_goods_num;
    this.stat.todoNum = res.data.stat.todo_number;
  },
}
</script>

<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 20px;
    padding-left: 10px;
    padding-right: 10px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
    border-color: rgba(0, 0, 0, .05);

    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }

      .icon-people {
        background: #40c9c6;
      }

      .icon-message {
        background: #36a3f7;
      }

      .icon-money {
        background: #f4516c;
      }

      .icon-shopping {
        background: #34bfa3
      }
    }

    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 6px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px 20px 26px 0px;
      margin-left: 0px;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}

@media (max-width:550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    .svg-icon {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>
