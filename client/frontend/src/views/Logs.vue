<template>
  <div class="logs">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>日志查看</span>
          <div class="header-actions">
            <el-select v-model="logLevel" placeholder="日志级别" style="width: 120px; margin-right: 10px;">
              <el-option label="全部" value="all" />
              <el-option label="调试" value="debug" />
              <el-option label="信息" value="info" />
              <el-option label="警告" value="warn" />
              <el-option label="错误" value="error" />
            </el-select>
            <el-button @click="refreshLogs">
              <el-icon><Refresh /></el-icon>
              刷新
            </el-button>
            <el-button @click="clearLogs">
              <el-icon><Delete /></el-icon>
              清空
            </el-button>
          </div>
        </div>
      </template>
      
      <div class="log-container">
        <div 
          v-for="(log, index) in filteredLogs" 
          :key="index" 
          class="log-item"
          :class="getLogClass(log.level)"
        >
          <span class="log-time">{{ log.time }}</span>
          <span class="log-level">{{ log.level.toUpperCase() }}</span>
          <span class="log-message">{{ log.message }}</span>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh, Delete } from '@element-plus/icons-vue'

const logs = ref([])
const logLevel = ref('all')
const autoRefresh = ref(true)
let refreshTimer = null

const filteredLogs = computed(() => {
  if (logLevel.value === 'all') {
    return logs.value
  }
  return logs.value.filter(log => log.level === logLevel.value)
})

const getLogClass = (level) => {
  return `log-${level}`
}

const refreshLogs = () => {
  // TODO: 从后端获取日志
  generateMockLogs()
  ElMessage.success('日志已刷新')
}

const clearLogs = () => {
  logs.value = []
  ElMessage.success('日志已清空')
}

const generateMockLogs = () => {
  const levels = ['debug', 'info', 'warn', 'error']
  const messages = [
    '应用启动成功',
    '正在连接到服务器...',
    '连接建立成功',
    '收到心跳包',
    '发送数据包',
    '连接异常断开',
    '重连中...',
    '连接恢复成功'
  ]
  
  const newLogs = []
  for (let i = 0; i < 20; i++) {
    const level = levels[Math.floor(Math.random() * levels.length)]
    const message = messages[Math.floor(Math.random() * messages.length)]
    newLogs.push({
      time: new Date().toLocaleTimeString(),
      level: level,
      message: message
    })
  }
  
  logs.value = [...newLogs, ...logs.value].slice(0, 100) // 最多保留100条日志
}

const startAutoRefresh = () => {
  if (autoRefresh.value) {
    refreshTimer = setInterval(() => {
      generateMockLogs()
    }, 5000) // 每5秒刷新一次
  }
}

const stopAutoRefresh = () => {
  if (refreshTimer) {
    clearInterval(refreshTimer)
    refreshTimer = null
  }
}

onMounted(() => {
  refreshLogs()
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<style scoped>
.logs {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-actions {
  display: flex;
  align-items: center;
}

.log-container {
  height: 500px;
  overflow-y: auto;
  background-color: #1e1e1e;
  color: #ffffff;
  padding: 10px;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 12px;
  line-height: 1.4;
}

.log-item {
  display: flex;
  margin-bottom: 2px;
  padding: 2px 0;
}

.log-time {
  color: #888;
  margin-right: 10px;
  min-width: 80px;
}

.log-level {
  margin-right: 10px;
  min-width: 50px;
  font-weight: bold;
}

.log-message {
  flex: 1;
}

.log-debug {
  color: #9cdcfe;
}

.log-info {
  color: #4fc1ff;
}

.log-warn {
  color: #ffcc02;
}

.log-error {
  color: #f44747;
}

.log-container::-webkit-scrollbar {
  width: 8px;
}

.log-container::-webkit-scrollbar-track {
  background: #2d2d2d;
}

.log-container::-webkit-scrollbar-thumb {
  background: #555;
  border-radius: 4px;
}

.log-container::-webkit-scrollbar-thumb:hover {
  background: #777;
}
</style>
