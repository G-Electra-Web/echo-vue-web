<script setup lang="ts">
import { ref } from "vue";
import {
    CircleUser,
    Home,
    Menu,
    Package2,
    Users,
    Presentation,
    Clock,
} from "lucide-vue-next";

import { Badge } from "@/components/ui/badge";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import Input from "@/components/ui/input/Input.vue";
import Card from "@/components/ui/card/Card.vue";
const activeSection = ref("dashboard");

const setActiveSection = (section: string) => {
    activeSection.value = section;
};
</script>

<template>
    <div
        class="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]"
    >
        <div class="hidden border-r bg-muted/40 md:block">
            <div class="flex h-full max-h-screen flex-col gap-2">
                <div
                    class="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6"
                >
                    <a href="/" class="flex items-center gap-2 font-semibold">
                        <Package2 class="h-6 w-6" />
                        <span class="">G - Electra | Admin</span>
                    </a>
                </div>
                <div class="flex-1">
                    <nav
                        class="grid items-start px-2 text-sm font-medium lg:px-4"
                    >
                        <a
                            href="#"
                            @click="setActiveSection('dashboard')"
                            :class="[
                                'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                activeSection === 'dashboard'
                                    ? 'bg-muted text-primary'
                                    : 'text-muted-foreground',
                            ]"
                            class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                        >
                            <Home class="h-4 w-4" />
                            Projects List
                        </a>
                        <a
                            href="#"
                            @click.prevent="setActiveSection('noticeBoard')"
                            :class="[
                                'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                activeSection === 'noticeBoard'
                                    ? 'bg-muted text-primary'
                                    : 'text-muted-foreground',
                            ]"
                            class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                        >
                            <Presentation class="h-4 w-4" />
                            Notice Board
                            <Badge
                                class="ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full"
                            >
                                6
                            </Badge>
                        </a>
                        <a
                            href="#"
                            @click.prevent="setActiveSection('members')"
                            :class="[
                                'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                activeSection === 'members'
                                    ? 'bg-muted text-primary'
                                    : 'text-muted-foreground',
                            ]"
                            class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                        >
                            <Users class="h-4 w-4" />
                            Memebers List
                        </a>
                        <a
                            href="#"
                            @click.prevent="setActiveSection('whitelisting')"
                            :class="[
                                'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                activeSection === 'whitelisting'
                                    ? 'bg-muted text-primary'
                                    : 'text-muted-foreground',
                            ]"
                            class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                        >
                            <Users class="h-4 w-4" />
                            Whitelisting
                        </a>
                        <a
                            href="#"
                            @click.prevent="setActiveSection('event')"
                            :class="[
                                'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                activeSection === 'event'
                                    ? 'bg-muted text-primary'
                                    : 'text-muted-foreground',
                            ]"
                            class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                        >
                            <Users class="h-4 w-4" />
                            Event
                        </a>
                        <a
                            href="#"
                            @click.prevent="setActiveSection('forms')"
                            :class="[
                                'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                activeSection === 'forms'
                                    ? 'bg-muted text-primary'
                                    : 'text-muted-foreground',
                            ]"
                            class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                        >
                            <Users class="h-4 w-4" />
                            Forms
                        </a>
                    </nav>
                </div>
            </div>
        </div>
        <div class="flex flex-col">
            <header
                class="flex h-14 items-center gap-4 border-b bg-muted/40 px-4 lg:h-[60px] lg:px-6"
            >
                <Sheet>
                    <SheetTrigger as-child>
                        <Button
                            variant="outline"
                            size="icon"
                            class="shrink-0 md:hidden"
                        >
                            <Menu class="h-5 w-5" />
                            <span class="sr-only">Toggle navigation menu</span>
                        </Button>
                    </SheetTrigger>
                    <SheetContent side="left" class="flex flex-col">
                        <nav class="grid gap-2 text-lg font-medium">
                            <a
                                href="#"
                                @click="setActiveSection('dashboard')"
                                :class="[
                                    'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                    activeSection === 'dashboard'
                                        ? 'bg-muted text-primary'
                                        : 'text-muted-foreground',
                                ]"
                                class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                            >
                                <Home class="h-4 w-4" />
                                Projects List
                            </a>
                            <a
                                href="#"
                                @click.prevent="setActiveSection('noticeBoard')"
                                :class="[
                                    'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                    activeSection === 'noticeBoard'
                                        ? 'bg-muted text-primary'
                                        : 'text-muted-foreground',
                                ]"
                                class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                            >
                                <Presentation class="h-4 w-4" />
                                Notice Board
                                <Badge
                                    class="ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full"
                                >
                                    6
                                </Badge>
                            </a>
                            <a
                                href="#"
                                @click.prevent="setActiveSection('members')"
                                :class="[
                                    'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                    activeSection === 'memebers'
                                        ? 'bg-muted text-primary'
                                        : 'text-muted-foreground',
                                ]"
                                class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                            >
                                <Users class="h-4 w-4" />
                                Memebers List
                            </a>
                            <a
                                href="#"
                                @click.prevent="
                                    setActiveSection('whitelisting')
                                "
                                :class="[
                                    'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                    activeSection === 'whitelisting'
                                        ? 'bg-muted text-primary'
                                        : 'text-muted-foreground',
                                ]"
                                class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                            >
                                <Users class="h-4 w-4" />
                                Whitelisting
                            </a>
                            <a
                                href="#"
                                @click.prevent="setActiveSection('event')"
                                :class="[
                                    'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                    activeSection === 'event'
                                        ? 'bg-muted text-primary'
                                        : 'text-muted-foreground',
                                ]"
                                class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                            >
                                <Users class="h-4 w-4" />
                                Event
                            </a>
                            <a
                                href="#"
                                @click.prevent="setActiveSection('forms')"
                                :class="[
                                    'flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary',
                                    activeSection === 'forms'
                                        ? 'bg-muted text-primary'
                                        : 'text-muted-foreground',
                                ]"
                                class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                            >
                                <Users class="h-4 w-4" />
                                Forms
                            </a>
                        </nav>
                        <div class="mt-auto"></div>
                    </SheetContent>
                </Sheet>
                <div class="w-full flex-1"></div>
                <DropdownMenu>
                    <DropdownMenuTrigger as-child>
                        <Button
                            variant="secondary"
                            size="icon"
                            class="rounded-full"
                        >
                            <CircleUser class="h-5 w-5" />
                            <span class="sr-only">Toggle user menu</span>
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                        <DropdownMenuLabel>My Account</DropdownMenuLabel>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem>Settings</DropdownMenuItem>
                        <DropdownMenuItem>Support</DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem>Logout</DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </header>
            <main class="flex flex-1 flex-col Agap-4 p-4 lg:gap-6 lg:p-6">
                <template v-if="activeSection === 'dashboard'">
                    <div class="flex items-center">
                        <h1 class="text-lg font-semibold md:text-2xl">
                            Project Requests
                        </h1>
                    </div>
                    <div
                        class="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm"
                    >
                        <div
                            class="flex flex-col items-center gap-1 text-center"
                        >
                            <h3 class="text-2xl font-bold tracking-tight">
                                <Clock class="h-10 w-10 ml-auto mr-auto mb-1" />
                                No Pending project Requests
                            </h3>
                        </div>
                    </div>
                </template>
                <template v-else-if="activeSection === 'noticeBoard'">
                    <div class="flex items-center">
                        <h1 class="text-lg font-semibold md:text-2xl">
                            Notice Board
                        </h1>
                    </div>
                    <div
                        class="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm"
                    >
                        <div
                            class="flex flex-col items-center gap-1 text-center"
                        >
                            <h3 class="text-2xl font-bold tracking-tight">
                                You have no projects
                            </h3>
                            <p class="text-sm text-muted-foreground">
                                You can start selling as soon as you add a
                                product.
                            </p>
                            <Button class="mt-4"> Create a project </Button>
                        </div>
                    </div>
                </template>
                <template v-if="activeSection === 'whitelisting'">
                    <ScrollArea class="h-[70%] w-full rounded-md border p-4">
                        <div class="flex items-center">
                            <h1 class="text-lg font-semibold md:text-2xl">
                                Emails Whitelisting
                            </h1>
                        </div>
                        <div
                            class="flex-col rounded-lg border border-dashed shadow-sm"
                        >
                            <div class="flex">
                                <Input
                                    class="w-[20rem] mt-6 ml-2"
                                    placeholder="Email"
                                />
                                <Button class="mt-6 ml-4"> Whitelist </Button>
                            </div>
                            <div class="mt-8 ml-4">
                                <h1 class="font-semibold">Users</h1>
                                <Table>
                                    <TableCaption
                                        >A list of Whitelisted
                                        Emails.</TableCaption
                                    >
                                    <TableHeader>
                                        <TableRow>
                                            <TableHead class="w-[100px]">
                                                SLNO
                                            </TableHead>
                                            <TableHead>Email</TableHead>
                                            <TableHead>Status</TableHead>
                                            <TableHead class="text-right">
                                                Button
                                            </TableHead>
                                        </TableRow>
                                    </TableHeader>
                                    <TableBody>
                                        <TableRow>
                                            <TableCell class="font-medium">
                                                1
                                            </TableCell>
                                            <TableCell
                                                >hmadhadi@gitam.in</TableCell
                                            >
                                            <TableCell
                                                >Not signed up yet</TableCell
                                            >
                                            <TableCell class="text-right">
                                                <Dialog>
                                                    <DialogTrigger>
                                                        <Button>
                                                            Delete
                                                        </Button>
                                                    </DialogTrigger>
                                                    <DialogContent>
                                                        <DialogHeader>
                                                            <DialogTitle
                                                                >Delete
                                                                Email</DialogTitle
                                                            >
                                                            <DialogDescription>
                                                                Are you sure you
                                                                want to delete
                                                            </DialogDescription>
                                                        </DialogHeader>

                                                        <DialogFooter>
                                                            <Button
                                                                variant="destructive"
                                                            >
                                                                Yes
                                                            </Button>
                                                        </DialogFooter>
                                                    </DialogContent>
                                                </Dialog>
                                            </TableCell>
                                        </TableRow>
                                        <TableRow>
                                            <TableCell class="font-medium">
                                                1
                                            </TableCell>
                                            <TableCell
                                                >hmadhadi@gitam.in</TableCell
                                            >
                                            <TableCell
                                                >Not signed up yet</TableCell
                                            >
                                            <TableCell class="text-right">
                                                <Dialog>
                                                    <DialogTrigger>
                                                        <Button>
                                                            Delete
                                                        </Button>
                                                    </DialogTrigger>
                                                    <DialogContent>
                                                        <DialogHeader>
                                                            <DialogTitle
                                                                >Delete
                                                                Email</DialogTitle
                                                            >
                                                            <DialogDescription>
                                                                Are you sure you
                                                                want to delete
                                                            </DialogDescription>
                                                        </DialogHeader>

                                                        <DialogFooter>
                                                            <Button
                                                                variant="destructive"
                                                            >
                                                                Yes
                                                            </Button>
                                                        </DialogFooter>
                                                    </DialogContent>
                                                </Dialog>
                                            </TableCell>
                                        </TableRow>
                                        <TableRow>
                                            <TableCell class="font-medium">
                                                1
                                            </TableCell>
                                            <TableCell
                                                >hmadhadi@gitam.in</TableCell
                                            >
                                            <TableCell
                                                >Not signed up yet</TableCell
                                            >
                                            <TableCell class="text-right">
                                                <Dialog>
                                                    <DialogTrigger>
                                                        <Button>
                                                            Delete
                                                        </Button>
                                                    </DialogTrigger>
                                                    <DialogContent>
                                                        <DialogHeader>
                                                            <DialogTitle
                                                                >Delete
                                                                Email</DialogTitle
                                                            >
                                                            <DialogDescription>
                                                                Are you sure you
                                                                want to delete
                                                            </DialogDescription>
                                                        </DialogHeader>

                                                        <DialogFooter>
                                                            <Button
                                                                variant="destructive"
                                                            >
                                                                Yes
                                                            </Button>
                                                        </DialogFooter>
                                                    </DialogContent>
                                                </Dialog>
                                            </TableCell>
                                        </TableRow>
                                        <TableRow>
                                            <TableCell class="font-medium">
                                                1
                                            </TableCell>
                                            <TableCell
                                                >hmadhadi@gitam.in</TableCell
                                            >
                                            <TableCell
                                                >Not signed up yet</TableCell
                                            >
                                            <TableCell class="text-right">
                                                <Dialog>
                                                    <DialogTrigger>
                                                        <Button>
                                                            Delete
                                                        </Button>
                                                    </DialogTrigger>
                                                    <DialogContent>
                                                        <DialogHeader>
                                                            <DialogTitle
                                                                >Delete
                                                                Email</DialogTitle
                                                            >
                                                            <DialogDescription>
                                                                Are you sure you
                                                                want to delete
                                                            </DialogDescription>
                                                        </DialogHeader>

                                                        <DialogFooter>
                                                            <Button
                                                                variant="destructive"
                                                            >
                                                                Yes
                                                            </Button>
                                                        </DialogFooter>
                                                    </DialogContent>
                                                </Dialog>
                                            </TableCell>
                                        </TableRow>
                                        <TableRow>
                                            <TableCell class="font-medium">
                                                1
                                            </TableCell>
                                            <TableCell
                                                >hmadhadi@gitam.in</TableCell
                                            >
                                            <TableCell
                                                >Not signed up yet</TableCell
                                            >
                                            <TableCell class="text-right">
                                                <Dialog>
                                                    <DialogTrigger>
                                                        <Button>
                                                            Delete
                                                        </Button>
                                                    </DialogTrigger>
                                                    <DialogContent>
                                                        <DialogHeader>
                                                            <DialogTitle
                                                                >Delete
                                                                Email</DialogTitle
                                                            >
                                                            <DialogDescription>
                                                                Are you sure you
                                                                want to delete
                                                            </DialogDescription>
                                                        </DialogHeader>

                                                        <DialogFooter>
                                                            <Button
                                                                variant="destructive"
                                                            >
                                                                Yes
                                                            </Button>
                                                        </DialogFooter>
                                                    </DialogContent>
                                                </Dialog>
                                            </TableCell>
                                        </TableRow>
                                    </TableBody>
                                </Table>
                            </div>
                        </div>
                    </ScrollArea>
                </template>
                <template v-if="activeSection === 'members'">
                    <div class="flex items-center">
                        <h1 class="text-lg font-semibold md:text-2xl">
                            Club Members
                        </h1>
                    </div>
                    <div
                        class="flex flex-1 rounded-lg border border-dashed shadow-sm"
                    >
                        <div class="w-full shadow-sm flex-col items-center">
                            <Avatar class="h-20 w-20">
                                <AvatarImage
                                    src="https://github.com/radix-vue.png"
                                    alt="@radix-vue"
                                />
                                <AvatarFallback>CN</AvatarFallback>
                            </Avatar>
                            <h1>Name</h1>
                            <h2>Role</h2>
                        </div>
                    </div>
                </template>
            </main>
        </div>
    </div>
</template>
