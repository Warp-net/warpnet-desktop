<!-- License generated by licensor(https://github.com/Marvin9/licensor).

Warpnet - Decentralized Social Network
Copyright (C) 2025 Vadim Filin, https://github.com/Warp-net,
<github.com.mecdy@passmail.net>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

WarpNet is provided “as is” without warranty of any kind, either expressed or implied.
Use at your own risk. The maintainers shall not be liable for any damages or data loss
resulting from the use or misuse of this software.
-->
<template>
  <div id="app" class="flex container h-screen w-full">
    <div class="flex container h-full w-full">
      <SideNav />

      <div class="w-full md:w-1/2 h-full overflow-y-scroll no-scrollbar">
        <div class="px-5 py-3 border-lighter flex items-center justify-between">
          <h1 class="text-xl font-bold">Notifications</h1>
          <i class="fas fa-cog text-xl text-blue"></i>
        </div>

        <div class="flex flex-col">
          <div class="flex flex-row justify-evenly text-sm">
            <button
              @click="submit('All')"
              class="w-full text-dark font-bold border-b-2 p-1 md:px-3 md:py-4 hover:bg-lightblue sm:truncate"
              :class="`${this.mode === 'All' ? 'border-blue' : ''}`"
            >
              All
            </button>
            <button
              @click="submit('Mentions')"
              class="w-full text-dark font-bold border-b-2 p-1 md:px-3 md:py-4 hover:bg-lightblue sm:truncate"
              :class="`${this.mode === 'Mentions' ? 'border-blue' : ''}`"
            >
              Mentions
            </button>
          </div>
        </div>

        <!-- notifications -->
        <div v-if="mode === 'All'">
          <div
            v-if="!all || all.length === 0"
            class="flex flex-col items-center justify-center w-full pt-10"
          >
            <div class="w-1/2 flex flex-col items-center justify-center">
              <p class="font-bold text-lg">No notifications yet</p>
              <p class="text-sm text-dark">Wait until some get shown here.</p>
            </div>
          </div>

          <div v-for="notification in all" :key="notification.id">
            <div
              class="w-full p-2 pt-1 pb-1 md:p-4 md:pt-2 md:pb-2 border-b hover:bg-lightest flex"
            >
              <div class="w-full">
                <div class="flex flex-row mr-2 md:mr-4 pt-1 text-2xl">
                  <i
                    v-if="notification.type === 'Replied'"
                    class="pt-1 fas fa-comment text-blue"
                  ></i>
                  <i
                    v-if="notification.type === 'Liked'"
                    class="pt-1 fas fa-heart text-red-600"
                  ></i>
                  <i
                    v-if="notification.type === 'Retweeted'"
                    class="pt-1 fas fa-retweet text-green-500"
                  ></i>

                  <a :href="`#/${profile.id}`">
                    <img
                      :src="`${profile.avatar || 'default_profile.png'}`"
                      class="h-8 w-8 ml-2 rounded-full flex-none"
                    />
                  </a>
                </div>
                <div class="flex items-center w-full">
                  <p class="font-sm">{{ profile.username }}</p>
                  <p class="font-sm">{{ "@" + profile.id }}</p>
                  <p class="text-sm text-dark ml-auto">
                    {{ $filters.timeago(notification.created_at) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-if="mode === 'Mentions'">
          <div
            v-if="!mentions || mentions.length === 0"
            class="flex flex-col items-center justify-center w-full pt-10"
          >
            <div class="w-1/2 flex flex-col items-center justify-center">
              <p class="font-bold text-lg">No notifications yet</p>
              <p class="text-sm text-dark">Wait until some get shown here.</p>
            </div>
          </div>
          <div v-for="notification in mentions" :key="notification.id">
            <div
              class="w-full p-2 pt-1 pb-1 md:p-4 md:pt-2 md:pb-2 border-b hover:bg-lightest flex"
            >
              <div class="w-full">
                <div class="flex flex-row mr-2 md:mr-4 pt-1 text-2xl">
                  <a :href="`#/${profile.id}`">
                    <img
                      :src="`${profile.avatar || 'default_profile.png'}`"
                      class="h-8 w-8 ml-2 rounded-full flex-none"
                    />
                  </a>
                </div>
                <div class="flex items-center w-full">
                  <p class="font-sm">{{ profile.username }}</p>
                  <p class="font-sm">{{ "@" + profile.id }}</p>
                  <p class="text-sm text-dark ml-auto">
                    {{ $filters.timeago(notification.created_at) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- default right bar -->
      <DefaultRightBar />
    </div>
  </div>
</template>

<script>
import SideNav from "../components/SideNav.vue";
import DefaultRightBar from "../components/DefaultRightBar.vue";

export default {
  name: "Notifications",
  components: {
    SideNav,
    DefaultRightBar,
  },
  data() {
    return {
      loading: false,
      mode: this.$route.query.m || "All",
    };
  },
  methods: {
    gotoHome() {
      this.$router.push({
        name: "Home",
      });
    },
    submit(m = this.mode) {
      this.mode = m;
      this.$router.replace({
        name: "Notifications",
        query: { m: m, hash: Date.now() },
      });
    },
  },
  async created() {
    console.log("loading component:", this.$options.name);

  },
};
</script>
