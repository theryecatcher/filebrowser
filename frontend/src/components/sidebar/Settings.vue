<template>
  <div
    v-for="setting in settings"
    :key="setting.id + '-sidebar'"
    :id="setting.id + '-sidebar'"
    class="card clickable"
    @click="setView(setting.id + '-main')"
    :class="{ hidden: !shouldShow(setting), 'active-settings': active(setting.id + '-main') }"
  >
    <div v-if="shouldShow(setting)" class="card-wrapper">{{ setting.label }}</div>
  </div>
</template>

<script>
import { state, getters, mutations } from "@/store";
import { settings } from "@/utils/constants";
import { router } from "@/router";

export default {
  name: "SidebarSettings",
  data() {
    return {
      settings, // Initialize the settings array in data
    };
  },
  computed: {
    currentHash: () => getters.currentHash(),
  },
  methods: {
    shouldShow(setting) {
      const perm = setting?.perm || {};
      // Check if all keys in setting.perm exist in state.user.perm and have truthy values
      return Object.keys(perm).every((key) => state.user.perm[key]);
    },
    active: (view) => state.activeSettingsView === view,
    setView(view) {
      if (state.route.path != "/settings") {
        router.push({ path: "/settings", hash: "#" + view }, () => {});
      } else {
        mutations.setActiveSettingsView(view);
      }
    },
  },
};
</script>
<style>
.active-settings {
  font-weight: bold;
  /* border-color: white; */
  border-style: solid;
}
</style>
