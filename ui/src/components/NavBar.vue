<script setup lang="ts">
import { Sheet, SheetTrigger, SheetContent } from '../components/ui/sheet'
import { Button } from '../components/ui/button'
</script>

<template>
    <nav class="flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6 sticky bg-gray-950 top-0 text-white">
      <!-- Desktop Navigation -->
      <div
        class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6 w-[80%] justify-between mx-auto">
        <a href="#" class="flex items-center gap-2 text-3xl font-semibold md:text-base">
          <span>
            <router-link to="/">G-Electra | GITAM</router-link>
          </span>
        </a>
        <div>
          <router-link class="text-white transition-colors hover:text-gray-900 hover:bg-gray-100 p-2 rounded"
          to="/login">Log In
        </router-link>
        <router-link class="text-white transition-colors hover:text-gray-900 hover:bg-gray-100 p-2 rounded"
          to="/signup">Sign Up
        </router-link>
        </div>
      
      </div>
  
      <!-- Mobile Navigation Trigger -->
      <Sheet>
        <SheetTrigger as-child>
          <Button variant="outline" size="icon" class="shrink-0 md:hidden">
            <Menu class="h-5 w-5" />
            <span class="sr-only">Toggle navigation menu</span>
          </Button>
        </SheetTrigger>
        <SheetContent side="left">
          <nav class="grid gap-6 text-lg font-medium">
            <a href="#" class="flex items-center gap-2 text-lg font-semibold">
              <span>
                <router-link to="/">G-Electra</router-link>
              </span>
            </a>
            <router-link class="text-gray-700 transition-colors hover:text-gray-900" to="/transactions">Transactions
            </router-link>
          </nav>
        </SheetContent>
      </Sheet>
    </nav>
  </template>