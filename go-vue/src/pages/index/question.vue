<template>
	<view class="content">
		<input type="text" name="question" v-model="formData.question" placeholder="输入问题">
		<button @click="buttonClick">发送</button>
		<textarea type= "text" name="answer" v-model="datalist" cols="30" rows="100" placeholder="答案"></textarea>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				formData: {
					question:''
				},
				datalist: '',// 定义绑定在页面上的data数据
			}
		},
		onLoad() {
			// 页面启动的生命周期，这里编写页面加载时的逻辑
		},
		methods: {
			buttonClick: function () {
				console.log("按钮被点了")
				var formData = this.formData
				uni.request({
				    url: 'http://127.0.0.1:8888/v1/doctor/question', //仅为示例，并非真实接口地址。
				    data: {
				        text: 'uni.request'
				    },
				    method:"POST",
					data:{
						question:formData.question
					},
					header: {
						'content-type': 'application/x-www-form-urlencoded',
					},
				    success: (res) => {
						console.log(res);
				        console.log(res.data.Answer);
				        this.text = 'request success';
						this.datalist = res.data.Answer;
				    }
				});
			},
		}
	}
</script>

<style>
	body {
	  font-family: 'Arial', sans-serif; /* 设置字体 */
	  background-color: #f4f4f4; /* 设置背景颜色 */
	  padding: 20px; /* 给body添加内边距 */
	}

	.content {
	  background-color: #fff; /* 容器背景色 */
	  max-width: 600px; /* 容器最大宽度 */
	  margin: 0 auto; /* 居中显示 */
	  padding: 20px; /* 容器内边距 */
	  border-radius: 8px; /* 边框圆角 */
	  box-shadow: 0 2px 4px rgba(0,0,0,0.1); /* 轻微的阴影效果 */
	}

	input[type="text"],
	textarea {
	  width: 100%; /* 宽度占满容器 */
	  padding: 10px; /* 输入框内边距 */
	  margin-bottom: 10px; /* 输入框与按钮的间距 */
	  border: 1px solid #ccc; /* 边框颜色 */
	  border-radius: 4px; /* 输入框边框圆角 */
	  box-sizing: border-box; /* 边框计算在宽度内 */
	}

	input[type="text"]:focus,
	textarea:focus {
	  outline: none; /* 移除默认的聚焦效果 */
	  border-color: #007bff; /* 聚焦时边框颜色 */
	  box-shadow: 0 0 8px rgba(0,123,255,0.2); /* 聚焦时的阴影效果 */
	}

	button {
	  width: 100%; /* 按钮宽度占满容器 */
	  padding: 10px; /* 按钮内边距 */
	  background-color: #007bff; /* 按钮背景色 */
	  color: white; /* 按钮文字颜色 */
	  border: none; /* 无边框 */
	  border-radius: 4px; /* 按钮边框圆角 */
	  cursor: pointer; /* 鼠标悬停时显示手形图标 */
	  font-size: 16px; /* 按钮文字大小 */
	  transition: background-color 0.3s; /* 背景色渐变效果 */
	}

	button:hover {
	  background-color: #0056b3; /* 鼠标悬停时按钮背景色变深 */
	}

	textarea {
	  resize: none; /* 禁止调整大小 */
	}
</style>