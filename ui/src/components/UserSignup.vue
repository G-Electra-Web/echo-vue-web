  <script setup lang="ts">
  import { ref } from 'vue';
  import { Button } from '../components/ui/button'
  import { useToast } from '../components/ui/toast/use-toast'
  import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '../components/ui/card'
  import { Input } from '../components/ui/input'
  import { Label } from '../components/ui/label'
  import { useRouter } from 'vue-router';
  import Toast from './ui/toast/Toast.vue';
  import Toaster from './ui/toast/Toaster.vue';


  const firstName = ref('');
  const lastName = ref('');
  const email = ref('');
  const regNum = ref('');
  const password = ref('');

  const router = useRouter();
  const { toast } = useToast();

  const handleSubmit = () => {
    // Basic validation
    if (!firstName.value || !lastName.value || !email.value || !regNum.value || !password.value) {
      alert('Please fill in all fields');
      return;
    }

    // Prepare data for submission
    const formData = {
      full_Name: firstName.value + " " + lastName.value,
      email: email.value,
      registration_number: regNum.value,
      password: password.value,
      membership_level: "Member",
    };

    // Example: Submit the formData to your backend API endpoint
    fetch('/api/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(data => {
      console.log('Success:', data);
      // Optionally, redirect to another page upon successful signup
      router.push('/signin');
    })
    .catch((error) => {
      console.error('Error:', error);
      toast({
        title: 'Failed to create account',
        description: 'Please try again later.',
        variant: 'destructive',
        action: {
          altText: 'Try again',
          callback: () => {
            handleSubmit(); // Retry submission on action click
          },
        },
      });
    });
  };

  </script>

  <template>
    <Toaster/>
    <Card class="mx-auto max-w-sm ">
      <CardHeader>
        <CardTitle class="text-xl">
          Sign Up
        </CardTitle>
        <CardDescription>
          Enter your information to create an account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="handleSubmit">
            <div class="grid gap-4">
          <div class="grid grid-cols-2 gap-4">
            <div class="grid gap-2">
              <Label for="firstname">First name</Label>
              <Input v-model="firstName" id="firstname" placeholder="Max" required />
            </div>
            <div class="grid gap-2">
              <Label for="lastname">Last name</Label>
              <Input v-model="lastName" id="lastname" placeholder="Robinson" required />
            </div>
          </div>
          <div class="grid gap-2">
            <Label for="email">Email</Label>
            <Input
              v-model="email"
              id="email"
              type="email"
              placeholder="m@example.com"
              required
            />
          </div>
          <div class="grid gap-2">
            <Label for="regNum">Gitam Registration Number</Label>
            <Input
              v-model="regNum"
              id="regNum"
              type="text"
              required
            />
          </div>
          <div class="grid gap-2">
            <Label for="password">Password</Label>
            <Input id="password" v-model="password" type="password" />
          </div>
          <Button type="submit" class="w-full">
            Create an account
          </Button>
        </div>
        </form>
      
        <div class="mt-4 text-center text-sm">
          Already have an account?
          <router-link to="/signin" class="underline">
            Sign in
          </router-link>
        </div>
      </CardContent>
    </Card>
  </template>