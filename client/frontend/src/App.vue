<template>
  <div id="app">
    <el-container>
      <el-header>
        <div class="header">
          <h1>HiTryRemote 客户端</h1>
          <div class="status">
            <el-tag :type="statusType">{{ statusText }}</el-tag>
          </div>
        </div>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <el-menu
            :default-active="activeMenu"
            class="sidebar-menu"
            @select="handleMenuSelect"
          >
            <el-menu-item index="dashboard">
              <el-icon><House /></el-icon>
              <span>仪表盘</span>
            </el-menu-item>
            <el-menu-item index="connections">
              <el-icon><Connection /></el-icon>
              <span>连接管理</span>
            </el-menu-item>
            <el-menu-item index="settings">
              <el-icon><Setting /></el-icon>
              <span>设置</span>
            </el-menu-item>
            <el-menu-item index="logs">
              <el-icon><Document /></el-icon>
              <span>日志</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { House, Connection, Setting, Document } from '@element-plus/icons-vue'

const router = useRouter()
const activeMenu = ref('dashboard')
const connectionStatus = ref('disconnected')

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

const handleMenuSelect = (index) => {
  activeMenu.value = index
  router.push(`/${index}`)
}

onMounted(() => {
  // 初始化应用
  console.log('HiTryRemote 客户端已启动')
})
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}

.header h1 {
  margin: 0;
  color: #409eff;
}

.status {
  display: flex;
  align-items: center;
}

.sidebar-menu {
  height: 100%;
  border-right: 1px solid #e6e6e6;
}

#app {
  height: 100vh;
}

.el-container {
  height: 100%;
}

.el-header {
  background-color: #f5f5f5;
  border-bottom: 1px solid #e6e6e6;
}

.el-aside {
  background-color: #fafafa;
}

.el-main {
  padding: 20px;
  background-color: #ffffff;
}
</style>
