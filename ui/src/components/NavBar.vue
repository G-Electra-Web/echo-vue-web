<script setup lang="ts">
import { Sheet, SheetTrigger, SheetContent } from '../components/ui/sheet'
import { Button } from '../components/ui/button'
import { useColorMode } from '@vueuse/core'
import { Icon } from '@iconify/vue'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '../components/ui/dropdown-menu'

const mode = useColorMode()
</script>

<template>
    <nav class="flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6 sticky bg-transparant top-0">
      <!-- Desktop Navigation -->
      <div
        class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6 w-[80%] justify-between mx-auto">
        <a href="#" class="flex items-center gap-2 text-3xl font-semibold md:text-base">
          <span>
            <router-link to="/">G-Electra | GITAM</router-link>
          </span>
        </a>
        <div>
          <router-link class="ransition-colors hover:text-gray-900 hover:bg-gray-100 p-2 rounded"
          to="/signin">Sign In
        </router-link>
        <router-link class="transition-colors hover:text-gray-900 hover:bg-gray-100 p-2 rounded"
          to="/signup">Sign Up
        </router-link>
          <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button variant="outline">
              <Icon icon="radix-icons:moon" class="rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
              <Icon icon="radix-icons:sun" class="absolute rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
              <span class="sr-only">Toggle theme</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem @click="mode = 'light'">
              Light
            </DropdownMenuItem>
            <DropdownMenuItem @click="mode = 'dark'">
              Dark
            </DropdownMenuItem>
            <DropdownMenuItem @click="mode = 'auto'">
              System
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
        </div>
      
      </div>
  
      <!-- Mobile Navigation Trigger -->
      <Sheet>
        <h1 class="md:hidden">G-Electra | GITAM</h1>
        <SheetTrigger as-child>
          <Button variant="outline" size="icon" class="shrink-0 md:hidden ml-auto">
            <Icon icon="radix-icons:hamburger-menu"/>
            <span class="sr-only">Toggle navigation menu</span>
          </Button>
        </SheetTrigger>
        <SheetContent side="right">
          <nav class="grid gap-6 text-lg font-medium">
            <a href="#" class="flex items-center gap-2 text-lg font-semibold">
              <span>
                <router-link to="/">G-Electra</router-link>
              </span>
            </a>
            <router-link class="transition-colors hover:text-gray-900" to="/signin">Sign In
            </router-link>
            <router-link class="transition-colors hover:text-gray-900" to="/signup">Sign Up
            </router-link>
            <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button variant="outline">
             
              <span>Toggle theme</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem @click="mode = 'light'">
              Light
            </DropdownMenuItem>
            <DropdownMenuItem @click="mode = 'dark'">
              Dark
            </DropdownMenuItem>
            <DropdownMenuItem @click="mode = 'auto'">
              System
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
          </nav>
        </SheetContent>
      </Sheet>
    </nav>
  </template>