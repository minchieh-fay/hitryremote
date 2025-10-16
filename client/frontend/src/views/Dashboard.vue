<template>
  <div class="dashboard">
    <el-row :gutter="20">
      <el-col :span="24">
        <el-card class="welcome-card">
          <h2>欢迎使用 HiTryRemote 客户端</h2>
          <p>这是一个基于 QUIC 协议的高性能代理客户端</p>
        </el-card>
      </el-col>
    </el-row>
    
    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="8">
        <el-card class="status-card">
          <template #header>
            <span>连接状态</span>
          </template>
          <div class="status-content">
            <el-tag :type="statusType" size="large">{{ statusText }}</el-tag>
            <p v-if="statusText === '未连接'">请先连接到服务器</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="8">
        <el-card class="stats-card">
          <template #header>
            <span>统计信息</span>
          </template>
          <div class="stats-content">
            <div class="stat-item">
              <span>连接数:</span>
              <span>{{ stats.connections }}</span>
            </div>
            <div class="stat-item">
              <span>传输量:</span>
              <span>{{ stats.transfer }}</span>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="8">
        <el-card class="actions-card">
          <template #header>
            <span>快速操作</span>
          </template>
          <div class="actions-content">
            <el-button 
              type="primary" 
              @click="connectToServer"
              :loading="connecting"
            >
              连接服务器
            </el-button>
            <el-button 
              type="danger" 
              @click="disconnectFromServer"
              :disabled="statusText === '未连接'"
            >
              断开连接
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

const connectionStatus = ref('disconnected')
const connecting = ref(false)
const stats = ref({
  connections: 0,
  transfer: '0 MB'
})

const statusType = computed(() => {
  switch (connectionStatus.value) {
    case 'connected':
      return 'success'
    case 'connecting':
      return 'warning'
    case 'error':
      return 'danger'
    default:
      return 'info'
  }
})

const statusText = computed(() => {
  switch (connectionStatus.value) {
    case 'connected':
      return '已连接'
    case 'connecting':
      return '连接中'
    case 'error':
      return '连接错误'
    default:
      return '未连接'
  }
})

const connectToServer = async () => {
  connecting.value = true
  connectionStatus.value = 'connecting'
  
  try {
    // TODO: 调用后端连接方法
    await new Promise(resolve => setTimeout(resolve, 2000)) // 模拟连接
    connectionStatus.value = 'connected'
    ElMessage.success('连接成功')
  } catch (error) {
    connectionStatus.value = 'error'
    ElMessage.error('连接失败: ' + error.message)
  } finally {
    connecting.value = false
  }
}

const disconnectFromServer = async () => {
  try {
    // TODO: 调用后端断开方法
    connectionStatus.value = 'disconnected'
    ElMessage.success('已断开连接')
  } catch (error) {
    ElMessage.error('断开连接失败: ' + error.message)
  }
}

onMounted(() => {
  // 初始化数据
  console.log('Dashboard 组件已加载')
})
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.welcome-card {
  text-align: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.welcome-card h2 {
  margin: 0 0 10px 0;
  font-size: 28px;
}

.welcome-card p {
  margin: 0;
  font-size: 16px;
  opacity: 0.9;
}

.status-content {
  text-align: center;
  padding: 20px 0;
}

.status-content p {
  margin-top: 10px;
  color: #666;
}

.stats-content {
  padding: 10px 0;
}

.stat-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
  padding: 5px 0;
  border-bottom: 1px solid #f0f0f0;
}

.actions-content {
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding: 10px 0;
}
</style>
