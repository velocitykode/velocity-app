import { Button, FormField, Input } from '@/components/ui';
import { usePageTitle } from '@/hooks/use-page-title';
import AuthLayout from '@/layouts/auth-layout';
import { Link, useForm } from '@inertiajs/react';
import { ArrowRight, Eye, EyeOff } from 'lucide-react';
import { useState } from 'react';

interface LoginProps {
  status?: string;
  canResetPassword?: boolean;
  errors?: {
    email?: string;
    password?: string;
  };
  old?: {
    email?: string;
  };
}

export default function Login({
  status,
  canResetPassword = false,
  errors = {},
  old = {},
}: LoginProps) {
  usePageTitle('Sign in');
  const [showPassword, setShowPassword] = useState(false);

  const { data, setData, post, processing } = useForm({
    email: old.email || '',
    password: '',
    remember: false,
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    post('/login');
  };

  return (
    <AuthLayout title="Welcome back" description="Sign in to your account to continue">
      {status && (
        <div className="mb-4 lg:mb-6 rounded-xl bg-emerald-50 p-3 lg:p-4 text-sm text-emerald-600 dark:bg-emerald-950/50 dark:text-emerald-400">
          {status}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4 lg:space-y-5">
        <FormField label="Email" htmlFor="email" error={errors.email}>
          <Input
            id="email"
            type="email"
            value={data.email}
            onChange={(e) => setData('email', e.target.value)}
            placeholder="you@example.com"
            autoComplete="email"
            autoFocus
            required
            error={!!errors.email}
          />
        </FormField>

        <FormField
          label="Password"
          htmlFor="password"
          error={errors.password}
        >
          <div className="relative">
            <Input
              id="password"
              type={showPassword ? 'text' : 'password'}
              value={data.password}
              onChange={(e) => setData('password', e.target.value)}
              autoComplete="current-password"
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
          {canResetPassword && (
            <Link
              href="/password/request"
              className="mt-1.5 block text-right text-xs text-[#1e3a8a] hover:underline dark:text-blue-400"
            >
              Forgot password?
            </Link>
          )}
        </FormField>

        <label className="flex cursor-pointer items-center gap-2.5">
          <input
            type="checkbox"
            checked={data.remember}
            onChange={(e) => setData('remember', e.target.checked)}
            className="h-4 w-4 rounded border-slate-300 text-[#1e3a8a] focus:ring-[#1e3a8a]/20 dark:border-zinc-600 dark:bg-zinc-800"
          />
          <span className="text-sm text-slate-600 dark:text-slate-400">Remember me</span>
        </label>

        <Button
          type="submit"
          fullWidth
          size="lg"
          loading={processing}
          iconRight={ArrowRight}
        >
          Sign in
        </Button>

        <p className="text-center text-sm text-slate-500 dark:text-slate-400">
          New to Velocity?{' '}
          <Link href="/register" className="font-medium text-[#1e3a8a] hover:underline dark:text-blue-400">
            Create an account
          </Link>
        </p>
      </form>
    </AuthLayout>
  );
}
