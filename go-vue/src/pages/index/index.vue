<template>
  <view>
    <button @click="goComment">去评论列表</button>
  </view>
</template>

<script>
export default {
  methods: {
    goComment() {
      uni.navigateTo({ url: '/pages/data/comment' })
    }
  }
}
</script>

<style>
	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.logo {
		height: 200rpx;
		width: 200rpx;
		margin: 200rpx auto 50rpx auto;
	}

	.text-area {
		display: flex;
		justify-content: center;
	}

	.title {
		font-size: 36rpx;
		color: #8f8f94;
	}
</style>
