
<template>
  <div class="container mt-4">
    <div class="table-responsive">
      <table class="table table-striped table-bordered">
        <thead class="thead-dark">
          <tr>
			 <th scope="col">ID</th>
            <th scope="col">设备ID</th>
            <th scope="col">用户ID</th>
            <th scope="col">数据记录时间</th>
            <th scope="col">心率</th>
            <th scope="col">收缩压</th>
            <th scope="col">舒张压</th>
            <th scope="col">体温</th>
            <th scope="col">步数</th>
            <th scope="col">睡眠时长</th>
            <th scope="col">活动卡路里</th>
            <th scope="col">血糖浓度</th>
            <th scope="col">体重</th>
            <th scope="col">身高</th>
            <th scope="col">设备状态</th>
            <th scope="col">操作</th>
			
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in dataList" :key="index">
			  <td>{{item.Id}}</td>
            <td>{{ item.device_id }}</td>
            <td>{{ item.user_id }}</td>
            <td>{{ item.timestamp }}</td>
            <td>{{ item.heart_rate }}</td>
            <td>{{ item.blood_pressure_systolic }}</td>
            <td>{{ item.blood_pressure_diastolic }}</td>
            <td>{{ item.body_temperature }}</td>
            <td>{{ item.steps }}</td>
            <td>{{ item.sleep_duration_minutes }}</td>
            <td>{{ item.activity_calories_burned }}</td>
            <td>{{ item.blood_glucose }}</td>
            <td>{{ item.weight }}</td>
            <td>{{ item.height }}</td>
            <td>{{ item.device_status }}</td>
            <td>
              <button @click="confirmDelete(item.Id)" class="btn btn-danger btn-sm">删除</button>
			  
            </td>
			
          </tr>
        </tbody>
      </table>
    </div>
    <div v-if="loading" class="text-center mt-3">
      <p>Loading...</p>
    </div>
    <div v-if="message" class="alert alert-danger mt-3">
      {{ message }}
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dataList: [], // 初始化数据列表为空数组
      loading: false, // 加载状态
      message: '', // 消息提示
    };
  },
  methods: {
    getDataList() {
      this.loading = true; // 开始加载数据时显示加载状态
      uni.request({
        url: "http://127.0.0.1:8888/v1/data/list", // 获取列表的接口地址
        method: "GET",
        success: (res) => {
          console.log(res.data);
          if (res.data.message === '健康设备展示数据成功') {
            this.dataList = res.data.data; // 将数据正确赋值给dataList
            console.log(this.dataList); // 打印数据列表
          } else {
            this.message = res.data.message; // 显示错误消息
          }
          this.loading = false; // 结束加载时隐藏加载状态
        },
        fail: (error) => {
          this.message = "网络错误"; // 显示网络错误消息
          this.loading = false; // 结束加载时隐藏加载状态
        }
      });
    },
    confirmDelete(itemId) {
      if (confirm("确认要删除这条记录吗？")) { // 确认删除操作
        this.loading = true; // 开始删除时显示加载状态
        uni.request({
          url: `http://127.0.0.1:8888/v1/data/delete/${itemId}`, // 删除单个数据项的接口地址，根据传入的itemId删除
          method: "DELETE",
          success: (res) => {
            console.log(res.data);
            if (res.data.message === '删除成功') {
              // 删除成功后更新前端列表
              this.dataList = this.dataList.filter(item => item.id !== itemId);
              alert('删除成功');
            } else {
              this.message = res.data.message; // 显示错误消息
            }
            this.loading = false; // 结束删除时隐藏加载状态
          },
          fail: (error) => {
            this.message = "网络错误"; // 显示网络错误消息
            this.loading = false; // 结束删除时隐藏加载状态
          }
        });
      }
    }
  },
  mounted() {
    this.getDataList(); // 组件挂载时获取数据
  }
};
</script>

<style scoped>
.container {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
  font-family: 'Arial', sans-serif;
  background-color: #ffffff;
  border-radius: 8px;
  overflow: hidden;
}

th, td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #dddddd;
}

th {
  background-color: #007bff;
  color: #ffffff;
  font-weight: bold;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}

tr:hover {
  background-color: #e9f5ff;
}

@media screen and (max-width: 768px) {
  table, thead, tbody, th, td, tr {
    display: block;
  }

  th, td {
    padding: 10px;
    text-align: right;
    position: relative;
  }

  th::before, td::before {
    content: attr(data-label);
    position: absolute;
    left: 0;
    width: 50%;
    padding-left: 10px;
    font-weight: bold;
    text-align: left;
  }
}
</style>
