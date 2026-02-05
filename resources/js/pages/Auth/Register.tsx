import { Button, FormField, Input } from '@/components/ui';
import { usePageTitle } from '@/hooks/use-page-title';
import AuthLayout from '@/layouts/auth-layout';
import { Link, useForm } from '@inertiajs/react';
import { ArrowRight, Eye, EyeOff } from 'lucide-react';
import { useState } from 'react';

interface RegisterProps {
  errors?: {
    name?: string;
    email?: string;
    password?: string;
    password_confirmation?: string;
  };
  old?: {
    name?: string;
    email?: string;
  };
}

export default function Register({ errors = {}, old = {} }: RegisterProps) {
  usePageTitle('Create account');
  const [showPassword, setShowPassword] = useState(false);

  const { data, setData, post, processing } = useForm({
    name: old.name || '',
    email: old.email || '',
    password: '',
    password_confirmation: '',
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    post('/register');
  };

  return (
    <AuthLayout title="Create account" description="Get started with Velocity in seconds">
      <form onSubmit={handleSubmit} className="space-y-4 lg:space-y-5">
        <FormField label="Name" htmlFor="name" error={errors.name}>
          <Input
            id="name"
            type="text"
            value={data.name}
            onChange={(e) => setData('name', e.target.value)}
            placeholder="Your name"
            autoComplete="name"
            autoFocus
            required
            error={!!errors.name}
          />
        </FormField>

        <FormField label="Email" htmlFor="email" error={errors.email}>
          <Input
            id="email"
            type="email"
            value={data.email}
            onChange={(e) => setData('email', e.target.value)}
            placeholder="you@example.com"
            autoComplete="email"
            required
            error={!!errors.email}
          />
        </FormField>

        <FormField label="Password" htmlFor="password" error={errors.password}>
          <div className="relative">
            <Input
              id="password"
              type={showPassword ? 'text' : 'password'}
              value={data.password}
              onChange={(e) => setData('password', e.target.value)}
              autoComplete="new-password"
              required
              error={!!errors.password}
              className="pr-12"
            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 dark:hover:text-slate-300"
              tabIndex={-1}
            >
              {showPassword ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
            </button>
          </div>
        </FormField>

        <FormField
          label="Confirm Password"
          htmlFor="password_confirmation"
          error={errors.password_confirmation}
        >
          <Input
            id="password_confirmation"
            type={showPassword ? 'text' : 'password'}
            value={data.password_confirmation}
            onChange={(e) => setData('password_confirmation', e.target.value)}
            autoComplete="new-password"
            required
            error={!!errors.password_confirmation}
          />
        </FormField>

        <Button
          type="submit"
          fullWidth
          size="lg"
          loading={processing}
          iconRight={ArrowRight}
        >
          Create account
        </Button>

        <p className="text-center text-sm text-slate-500 dark:text-slate-400">
          Already have an account?{' '}
          <Link href="/login" className="font-medium text-[#1e3a8a] hover:underline dark:text-blue-400">
            Sign in
          </Link>
        </p>
      </form>
    </AuthLayout>
  );
}
