<template>
	<view class="">
		<table>
			<tr>
				<td>提问时间</td>
				<td>修改时间</td>
				<td>提问问题</td>
				<td>问题答案</td>
			</tr>
			<tr v-for="v in quesHisList">
				<td>{{v.CreatedAt}}</td>
				<td>{{v.UpdatedAt}}</td>
				<td>{{v.Question}}</td>
				<td>{{v.Answer}}</td>
			</tr>
		</table>
	</view>
</template>

<script>
	import { URLSearceParams } from 'url'
	export default {
				data() {
					return {
						quesHisList:[]
					};
				},
				methods: {
					getUserList() {
						uni.request({
						    url: "http://127.0.0.1:8888/v1/doctor/list", //仅为示例，并非真实接口地址。
						    data: {
						        //text: 'uni.request'
						    },
							method:"POST",
							header: {
								'content-type': 'application/x-www-form-urlencoded',
							},
						    success: (res) => {
						         console.log(res);
						        this.text = 'request success';
								if (res.statusCode == 200) {
									this.quesHisList = res.data.List
									console.log(res.data);
								} else {
									alert(res.data.message)
								}
						    },
							fail:(error)=> {
								alert("网络错误")
							}
						});
					}
				},
				mounted(){
					this.getUserList()
				}
			}
</script>

<style>
  body {
    font-family: 'Arial', sans-serif;
    background-color: #f4f4f4;
    margin: 0;
    padding: 20px;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    margin: 20px 0;
  }

  th, td {
    border: 1px solid #ddd;
    padding: 15px;
    text-align: left;
  }

  th {
    background-color: #4CAF50;
    color: white;
  }

  tr:nth-child(even) {
    background-color: #f2f2f2;
  }

  tr:hover {
    background-color: #ddd;
  }

  .view {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }
</style>
